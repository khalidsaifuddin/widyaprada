package conn

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/ProjectWidyaprada/backend/config"
)

type Cache struct {
	Pool *redis.Pool
}

type CacheService interface {
	Ping() error
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl int64) error
	Exists(key string) (bool, error)
	Delete(key string) error
}

func NewCacheService(pool *redis.Pool) CacheService {
	return &Cache{Pool: pool}
}

func CreateRedisPool(addr, password string, maxIdle int, dbIndex int) (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle: maxIdle,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, redis.DialDatabase(dbIndex))
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
	conn := pool.Get()
	defer conn.Close()
	if _, err := conn.Do("PING"); err != nil {
		return pool, err
	}
	return pool, nil
}

func (c *Cache) Ping() error {
	conn := c.Pool.Get()
	defer conn.Close()
	_, err := redis.String(conn.Do("PING"))
	return err
}

func (c *Cache) Get(key string) ([]byte, error) {
	conn := c.Pool.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func (c *Cache) Set(key string, value []byte, ttl int64) error {
	conn := c.Pool.Get()
	defer conn.Close()
	if _, err := conn.Do("SET", key, value); err != nil {
		return err
	}
	_, err := conn.Do("EXPIRE", key, ttl)
	return err
}

func (c *Cache) Exists(key string) (bool, error) {
	conn := c.Pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}

func (c *Cache) Delete(key string) error {
	conn := c.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}

func InitRedis(cfg config.Config) (CacheService, *redis.Pool) {
	addr := cfg.RedisHost + ":" + cfg.RedisPort
	pool, err := CreateRedisPool(addr, cfg.RedisPassword, cfg.RedisMaxIdle, cfg.RedisDBIndex)
	if err != nil {
		panic(fmt.Sprintf("redis: %v", err))
	}
	log.Print("Connected to Redis")
	return NewCacheService(pool), pool
}
