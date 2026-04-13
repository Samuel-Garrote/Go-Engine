package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx2 = context.Background()

func main11(){
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	client.RPush(ctx, "jobs", "job1", "job2", "job3")
	val,_ := client.LPop(ctx, "jobs").Result()
	fmt.Println(val)
	val,_ = client.LPop(ctx, "jobs").Result()
	fmt.Println(val)

blVal, err := client.BLPop(ctx, 0, "jobs").Result()
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println(blVal[1])
}