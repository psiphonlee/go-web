package internal

import (
	"gomap/global"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBBASE interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	// 将传入的字符串前缀和单复数形式参数应用到 GORM 的命名策略中，
	// 并禁用迁移过程中的外键约束，返回最终生成的 GORM 配置信息。
	config := &gorm.Config{
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,   // 表前缀，在表名前添加前缀，如添加用户模块的表前缀 user_
			SingularTable: singular, // 单复数形式, 如果为 true 则user 对应的表是 users
		},
		// 是否在迁移时禁用外键约束，默认false，表示会根据模型之间的关联自动生成外键约束语句
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)),
		logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      true,
		})

	var LogMode DBBASE
	switch global.EWA_CONFIG.App.DbType {
	case "mysql":
		LogMode = &global.EWA_CONFIG.MySQL // 这里的LogMode是MySQL结构体的指针
	default:
		LogMode = &global.EWA_CONFIG.MySQL
	}

	// 在config/gorm_mysql.go中,MySQL结构体实现了GetLogMode()方法
	switch LogMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
