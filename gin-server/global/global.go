package global

import (
	"gin-server/config/configModel"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	VIPER  *viper.Viper
	CONFIG configModel.Config
	LOG    *logrus.Logger
	TIME   time.Time
)
