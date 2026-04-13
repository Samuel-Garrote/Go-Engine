package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{}
var nextID = 1

func main2() {
	r := gin.Default()

	r.GET("/tasks/:id", func(c *gin.Context) {
		idStr:= c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.JSON(400, gin.H{"error" : "Invalid ID"})
		}
		for _, task := range tasks{
			if task.ID == id{
							c.JSON(200, task)
					return
			}
	}
	c.JSON(404, gin.H{"error" : "Task not found",
	})

	})

	r.POST("/tasks", func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		task.ID = nextID
		nextID++
		tasks = append(tasks, task)
		c.JSON(201, task)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "invalid id"})
        return
    }

    var updated Task
    if err := c.ShouldBindJSON(&updated); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    for i, task := range tasks {
        if task.ID == id {
            updated.ID = id
            tasks[i] = updated
            c.JSON(200, updated)
            return
        }
    }

    c.JSON(404, gin.H{"error": "task not found"})
})

r.DELETE("/tasks/:id", func(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "invalid id"})
        return
    }

    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            c.JSON(200, gin.H{"message": "task deleted"})
            return
        }
    }

    c.JSON(404, gin.H{"error": "task not found"})
})
	r.Run(":8080")
}