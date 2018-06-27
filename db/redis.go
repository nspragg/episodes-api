package db

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type Client struct {
	pool *redis.Pool
}

const (
	PROTOCAL = "tcp"
	SERVER   = "localhost:6379"
)

func Dial() (*Client, error) {
	client := new(Client)

	client.pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(PROTOCAL, SERVER)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	return client, nil
}

func compact(arr []string) []string {
	results := make([]string, 0)
	for _, v := range arr {
		if v != "" {
			results = append(results, v)
		}
	}
	return results
}

// Find ...
func (c *Client) Find(args ...interface{}) ([]string, error) {
	conn := c.pool.Get()
	defer conn.Close()

	raw, _ := redis.Strings(conn.Do("MGET", args...))
	return compact(raw), nil
}

// Close ..
func (c *Client) Close() error {
	return c.Close()
}
