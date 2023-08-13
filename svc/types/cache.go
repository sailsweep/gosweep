package types

import (
	"fmt"

	"github.com/go-redis/redis"
)

type CacheConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type Cache struct {
	Config CacheConfig
	Client interface{}
}

// implement the interface of service

func (c *Cache) Init() error {
	// connect to redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Config.Host, c.Config.Port),
		Password: c.Config.Pass,
		DB:       0,
	})

	// ping redis to check if it's connected
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	c.Client = client
	return nil
}

func (c *Cache) Close() (callback func(), err error) {
	return func() {
		err = c.Client.(*redis.Client).Close()
	}, err
}
