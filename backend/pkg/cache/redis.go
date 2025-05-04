package cache

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis(host string, port int, password string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		Password: password,
		DB:       db,
	})
	log.Println("Redis is ready")
}
func Redis() *redis.Client {
	return rdb
}
