package tool

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"log"
	"time"
)

type RedisStore struct {
	client  *redis.Client
	context context.Context
}

var RedisStoreObj *RedisStore

func InitRedisStore() *RedisStore {
	config := GetCofig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})
	RedisStoreObj = &RedisStore{
		client:  client,
		context: context.Background(),
	}

	base64Captcha.SetCustomStore(RedisStoreObj)
	return RedisStoreObj
}

func (s *RedisStore) Set(id string, value string) {
	err := s.client.Set(s.context, id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

func (s *RedisStore) Get(id string, clear bool) string {
	val, err := s.client.Get(s.context, id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := s.client.Del(s.context, id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}
