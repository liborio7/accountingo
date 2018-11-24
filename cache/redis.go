package cache

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
)

func Redis(o *Opt) Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     o.Addr,
		PoolSize: o.PoolSize,
	})
	return &redisCache{client}
}

type redisCache struct {
	client *redis.Client
}

func (c *redisCache) SetKey(m Model) error {
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

func (c *redisCache) GetKey(m Model) error {
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
