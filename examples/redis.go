package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)


var ctx3 = context.Background()

func main9(){
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err:= client.Ping(ctx).Err()
	if err != nil{
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis")


//***_______________________________
	err = client.Set(ctx, "name", "Go Engine", 10 * time.Second).Err()
	if err != nil{
		fmt.Println("Error:", err)
		return
	}

	value, err := client.Get(ctx, "name").Result()
	if err != nil{
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Value:" , value)
}