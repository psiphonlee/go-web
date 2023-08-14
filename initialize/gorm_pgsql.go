package initialize

import (
	"fmt"
	"gomap/initialize/internal"
	"gomap/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormPgSql 初始化 PostgreSQL 数据库
func GormPgSql() *gorm.DB {
	p := global.EWA_CONFIG.PgSQL

	if p.Dbname == "" {
		return nil
	}

	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,   // 是否使用简单协议
	}

	db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular))
	if err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)

		fmt.Println("====3-gorm====: gorm link postgresql success")
		return db
	}
}
