package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	prefix = "redis:%d"
	ttl    = time.Second * 30
)

func main() {
	//init redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	//params
	ctx := context.Background()
	key := fmt.Sprintf(prefix, 1)

	var u = User{
		Name: "s1ark",
		Age:  20,
	}

	//marshal struct to json
	val, _ := json.Marshal(u)

	//set cache with ttl(time to live)
	rdb.Set(ctx, key, val, ttl)

	//get cache
	str, _ := rdb.Get(ctx, key).Result()

	//unmarshal json to map
	var result map[string]interface{}
	json.Unmarshal([]byte(str), &result)
	fmt.Println(result)

	//del cache
	rdb.Del(ctx, key)

	//check is exist
	_, err := rdb.Get(ctx, key).Result()
	fmt.Println(err)
}
