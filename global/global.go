package global

import (
	"gomap/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	EWA_CONFIG *config.Configuration
	EWA_VIPER  *viper.Viper
	EWA_LOG    *zap.Logger
	EWA_DB     *gorm.DB
	EWA_REDIS  *redis.Client
)
