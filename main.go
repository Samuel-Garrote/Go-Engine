package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func heavyProcess(jobID string) string {
	var wg sync.WaitGroup
	results := make(chan string, 5)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(part int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			results <- fmt.Sprintf("%s part %d done", jobID, part)
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var summary string
	for message := range results {
		summary += message + " | "
	}
	return summary
}

func startWorker() {
	for {
		result, err := rdb.BLPop(ctx, 0, "Jobs_Queue").Result()
		if err != nil {
			fmt.Println("Worker error:", err)
			continue
		}

		jobID := result[1]
		fmt.Println("Processing:", jobID)

		rdb.Set(ctx, jobID, "processing", 0)
		summary := heavyProcess(jobID)
		fmt.Println("Result:", summary)
		rdb.Set(ctx, jobID, "done", 0)

		fmt.Println("Done:", jobID)
	}
}

func main() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	go startWorker()

	r := gin.Default()

	// healthcheck
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/jobs", func(c *gin.Context) {
		var payload map[string]string
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// validation
		if payload["message"] == "" {
			c.JSON(400, gin.H{"error": "message is required"})
			return
		}

		jobID := fmt.Sprintf("job:%d", time.Now().UnixNano())
		rdb.RPush(ctx, "Jobs_Queue", jobID)
		rdb.Set(ctx, jobID, "pending", 0)

		c.JSON(202, gin.H{
			"jobID":  jobID,
			"status": "pending",
		})
	})

	r.GET("/jobs/:id", func(c *gin.Context) {
		jobID := c.Param("id")
		val, err := rdb.Get(ctx, "job:"+jobID).Result()
		if err != nil {
			c.JSON(404, gin.H{"error": "job not found"})
			return
		}
		c.JSON(200, gin.H{
			"jobId":  jobID,
			"status": val,
		})
	})

	r.Run(":8080")
}