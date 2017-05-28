package main

import (
	"github.com/go-redis/redis"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client

}

func SetKey(pass Pass) {
	client := RedisClient()
	err := client.Set(pass.Id.String(), pass.Pass, pass.TTL).Err()
	if err != nil {
		panic(err)
	}
}

func GetKey(pass Pass) string {
	client := RedisClient()
	val, err := client.Get(pass.Id.String()).Result()
	if err != nil {
		panic(err)
	}

	return val
}

func DelKey(key string) {
	client := RedisClient()
	err := client.Del(key).Err()
	if err != nil {
		panic(err)
	}
}
