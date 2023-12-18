package cache

import (
	"context"
	"errors"
	"fmt"

	"github.com/isd-sgcu/oph66-backend/cfgldr"
	"github.com/redis/go-redis/v9"
)

func New(config *cfgldr.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.RedisConfig.Addr, config.RedisConfig.Port),
		Password: config.RedisConfig.Password,
	})

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.New("unable to ping redis")
	}

	return client, nil
}
