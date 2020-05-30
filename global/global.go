package global

import (
	"FuckingVersion1/config"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	oplogging "github.com/op/go-logging"
)

var (
	MyConfig config.Server
	MyVp     *viper.Viper
	MyDB 	 *gorm.DB
	MyLog    *oplogging.Logger
	MyRedis  *redis.Client
)

