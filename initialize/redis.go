package initialize

import (
	"FuckingVersion1/global"
	"github.com/go-redis/redis"
)

func Redis() {
	redisCfg := global.MyConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.MyLog.Error(err)
	} else {
		global.MyLog.Info("redis connect ping response:", pong)
		global.MyRedis = client
	}
}
