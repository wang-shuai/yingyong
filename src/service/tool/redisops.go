package tool

import (
	"time"
	"os"
	"syscall"
	"os/signal"
	"github.com/garyburd/redigo/redis"
)

func RedisInit(address string, maxIdle, timeout int) *redis.Pool{
	pool := newPool(address, maxIdle, timeout)
	cls(pool)
	return pool
}

func newPool(server string, maxIdle, timeout int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(timeout) * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cls(pool *redis.Pool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)

	go func() {
		<-c
		pool.Close()
		os.Exit(0)
	}()
}
