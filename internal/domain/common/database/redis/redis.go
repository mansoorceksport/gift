package redis

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

func newRedis(connectionString string, db int, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: password,
		DB:       db,
	})

	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	return client
}

type MutexRedis struct {
	redisClient *redis.Client
	mutex       *redsync.Redsync
}

func NewMutexRedis(connectionString string, db int, password string) *MutexRedis {
	redisClient := newRedis(connectionString, db, password)
	return &MutexRedis{
		redisClient: redisClient,
		mutex:       newRedisLock(redisClient),
	}
}

func newRedisLock(redisClient *redis.Client) *redsync.Redsync {
	pool := goredis.NewPool(redisClient)

	rs := redsync.New(pool)
	return rs
}
