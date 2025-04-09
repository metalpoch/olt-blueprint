package utils

import (
	"context"
	"encoding/json"

	core "github.com/metalpoch/olt-blueprint/core/model"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(url string) *redis.Client {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}

func SerializeModel(ld core.RedisDevice) (string, error) {
	jsonData, err := json.Marshal(ld)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func DeserializeModel(jsonData string) (core.RedisDevice, error) {
	var ld core.RedisDevice
	err := json.Unmarshal([]byte(jsonData), &ld)
	if err != nil {
		return core.RedisDevice{}, err
	}
	return ld, nil
}

func VerifyExistence(client *redis.Client, key string, ctx context.Context) bool {
	_, err := client.Exists(ctx, key).Result()
	return err != nil
}
