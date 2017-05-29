package main

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"time"
	"errors"
)

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client

}

func SetKey(pass Pass) uuid.UUID {
	client := RedisClient()
	pass.Id = uuid.New()
	pass.TTL = time.Minute * 2
	err := client.Set(pass.Id.String(), pass.Pass, pass.TTL).Err()
	if err != nil {
		panic(err)
	}
	return pass.Id
}

func GetKey(id string) (string, error) {
	client := RedisClient()
	val, err := client.Get(id).Result()
	if err != nil {
		return "", errors.New(err.Error())
	}

	return val, nil
}

func DelKey(key string) {
	client := RedisClient()
	err := client.Del(key).Err()
	if err != nil {
		panic(err)
	}
}
