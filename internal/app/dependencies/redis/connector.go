package redis

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/shitikovkirill/auth-service/internal/app/config"
)

var (
	client *Client
	once   = &sync.Once{}
)

// Load ...
func Load(cfg config.Redis) (err error) {
	once.Do(func() {
		cli := redis.NewClient(
			&redis.Options{
				Addr:     cfg.Address,
				Password: cfg.Password,
				PoolSize: cfg.PoolSize,
			})

		err = cli.Ping().Err()
		if err != nil {
			return
		}

		client = &Client{cli: cli}
	})

	return err
}

// GetRedis ...
func GetRedis() Cli {
	return client
}
