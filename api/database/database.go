package database

import (
	"context"
	"os"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(db int) *redis.Client {
	redisDatabase := redis.NewClient(&redis.Options{
		Addr :  os.Getenv("DATABASE"),
		Password : os.Getenv("DB_PASS"),
		DB :    db,
	})
	return redisDatabase
}