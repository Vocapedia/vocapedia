package cache

import (
	"fmt"
	"log"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.ReadValue().Redis.Host, config.ReadValue().Redis.Port),
		Password: config.ReadValue().Redis.Password,
		DB:       config.ReadValue().Redis.DB,
	})
	log.Println("Redis is ready")
}
func Redis() *redis.Client {
	return rdb
}
