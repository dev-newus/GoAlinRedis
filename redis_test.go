package GoAlinRedis

import (
	"context"
	"testing"
	"time"
)

var (
	host     = "redis-16293.c8.us-east-1-4.ec2.cloud.redislabs.com:16293"
	username = "default"
	password = "2YCpMDk9LnGAyrWVCy8AApBIAoZ6CjrW"
	database = 0
)

func TestConnectRedis(t *testing.T) {
	client := newClient(host, username, password, database)

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		t.Errorf("Error pinging Redis: %s", err)
	}
}

func TestSetRedis(t *testing.T) {
	client := newClient(host, username, password, database)

	asd := set(client, "test", "ini test", 24*time.Hour)
	if asd != true {
		t.Errorf("Error set to redis")
	}
}

func TestGetRedis(t *testing.T) {
	client := newClient(host, username, password, database)

	status, _ := get(client, "test")
	if status != true {
		t.Errorf("Error get from redis")
	}
}
