package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Job3 struct {
    ID     int    `json:"id"`
    Status string `json:"status"`
}

func main() {
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    job := Job{ID: 1, Status: "pending"}

    data, _ := json.Marshal(job)
    client.Set(ctx, "job:1", data, 0)

    val, _ := client.Get(ctx, "job:1").Result()
    
    var result Job3
    json.Unmarshal([]byte(val), &result)
    

    fmt.Println(result)
}