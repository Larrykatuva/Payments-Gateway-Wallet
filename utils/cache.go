package utils

import (
	"context"
	"encoding/json"
	"errors"
	"example.com/m/initializers"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	expiry time.Duration
	data   interface{}
	key    *string
	client *redis.Client
}

func NewCacheService() *Cache {
	return &Cache{
		expiry: 1 * time.Minute,
		client: initializers.RedisClient,
	}
}

func (c *Cache) WithExpiry(duration time.Duration) *Cache {
	c.expiry = duration
	return c
}

func (c *Cache) WithData(data interface{}) *Cache {
	c.data = data
	return c
}

func (c *Cache) WithKey(key string) *Cache {
	c.key = &key
	return c
}

func (c *Cache) Save(ctx context.Context) error {
	if c.key == nil {
		return errors.New("key to cache value must be set")
	}
	data, err := json.Marshal(c.data)
	if err != nil {
		return err
	}
	if err = c.client.Set(ctx, *c.key, data, c.expiry).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Cache) Clear(ctx context.Context) error {
	statusCmd := c.client.FlushAll(ctx)
	if err := statusCmd.Err(); err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetData(ctx context.Context) ([]byte, error) {
	if c.key == nil {
		return nil, errors.New("key to cache value must be set")
	}
	storedData, err := c.client.Get(ctx, *c.key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return storedData, nil
}
