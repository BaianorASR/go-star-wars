package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/BaianorASR/go-star-wars/database"
	"github.com/americanas-go/errors"
)

func Set(key string, value interface{}) (err error) {

	if value == nil || value == "" {
		println("Value is nil or empty")
		return errors.Internalf("Canot set nil or zero value to redis")
	}

	rdb := database.GetRedis()

	json, err := json.Marshal(value)
	if err != nil {
		return errors.New("Error marshalling value")
	}

	err = rdb.Set(context.Background(), key, json, 0).Err()
	if err != nil {
		return errors.Internalf("Error setting key %s", key)
	}
	return nil
}

func Get[V interface{}](key string) (*V, error) {
	var result V

	rdb := database.GetRedis()
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, nil
	}

	if val == "" {
		fmt.Println("Val ->", val)
		println("Cache miss")
		return nil, nil
	}

	if reflect.TypeOf(val).Name() == "string" && reflect.TypeOf(result).Name() == "string" {
		return nil, nil
	}

	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return nil, errors.New("Error unmarshalling value")
	}

	return &result, nil
}
