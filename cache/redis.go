package cache

import (
	"github.com/garyburd/redigo/redis"
)

type Cache interface {
	Get(key string) (interface{},error)
	Set(key string, value interface{}) error
	SetWithTTL(key string, value interface{}, ttl int) error
}

type redisCache struct {
	pool redis.Pool
}

func (r *redisCache) Get(key string) (interface{},error) {
	c := r.pool.Get()
	defer c.Close()

	value, err := redis.String(c.Do("GET", key))

	if err != nil {
		return nil,err
	}

	return value,nil
}

func (r *redisCache) Set(key string, value interface{}) error {
	c := r.pool.Get()
	defer c.Close()

	_, err := c.Do("SET", key, value)

	if err != nil {
		return err
	}

	return err
}

func (r *redisCache) SetWithTTL(key string, value interface{}, ttl int) error {
	c := r.pool.Get()
	defer c.Close()

	_,err := c.Do("SETEX", key,ttl,value)
	if err != nil {
		return  err
	}
	return err
}

func NewRedisCache(addr string, maxConnections int) Cache {
	pool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", addr)

		if err != nil {
			return nil, err
		}

		return c, err
	}, maxConnections)

	return &redisCache{
		pool: *pool,
	}
}
