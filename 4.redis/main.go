package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var u = User{
	Name: "s1ark",
	Age:  20,
}

const (
	ttl = time.Second * 30
)

func main() {
	//Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ctx := context.Background()
	key := "user"

	//Serialize struct to json
	val, _ := json.Marshal(u)

	//Store data in Redis with TTL(time to live)
	rdb.Set(ctx, key, val, ttl)

	str, _ := rdb.Get(ctx, key).Result()

	//Deserialize json to map
	var result map[string]interface{}
	json.Unmarshal([]byte(str), &result)
	fmt.Println(result)

	rdb.Del(ctx, key)

	//Verify whether the key still exists in Redis
	_, err := rdb.Get(ctx, key).Result()
	fmt.Println(err)
}
