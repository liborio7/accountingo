package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
)

type RedisOpt struct {
	Addr     string
	PoolSize int
}

func Redis(o *RedisOpt) *Service {
	client := redis.NewClient(&redis.Options{
		Addr:     o.Addr,
		PoolSize: o.PoolSize,
	})
	return &Service{&redisClient{client: client}}
}

type redisClient struct {
	client *redis.Client
}

func (c *redisClient) SetKey(_ context.Context, m Model) error {
	bytes, marshalErr := json.Marshal(m)
	if marshalErr != nil {
		return marshalErr
	}
	if m.GetKey() == "" {
		return errors.New("key is nil")
	}
	_, setErr := c.client.Set(m.GetKey(), bytes, -1).Result()
	if setErr != nil {
		return setErr
	}
	return nil
}

func (c *redisClient) GetKey(_ context.Context, m Model) error {
	bytes, getErr := c.client.Get(m.GetKey()).Result()
	if getErr != nil {
		return getErr
	}
	unmarshalErr := json.Unmarshal([]byte(bytes), m)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	return nil
}
