/**
 * Author: Diki Rahmad Sandi
 * File: redis.go
 */

package GoAlinRedis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// function for create new client
func newClient(host string, username string, password string, database int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       database,
	})
	return client
}

// function to set
func set(connection *redis.Client, key string, value string, expire time.Duration) bool {
	ctx := context.Background()
	err := connection.Set(ctx, key, value, expire).Err()
	if err != nil {
		return false
	}
	return true

}

// function to get
func get(connection *redis.Client, key string) (bool, string) {
	ctx := context.Background()
	val, err := connection.Get(ctx, key).Result()
	if err != nil {
		return false, "Failed"
	} else {
		return true, val
	}
}
