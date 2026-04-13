package main

import (
	"fmt"
	"time"
)
type Job struct{
	ID int
	Payload string
	Status string
}

func worker(queue chan Job){
for job := range queue{
		fmt.Println("Processing...", job.ID)
		time.Sleep(2 * time.Second)
		fmt.Println("Done Job", job.ID)
}
}

func main10(){
	queue := make(chan Job, 10)

	go worker(queue)
	queue <- Job{ID: 1, Payload: "first job", Status: "Pending"}
	queue <- Job{ID: 2, Payload: "second job", Status: "Pending"}
	queue <- Job{ID: 3, Payload: "third job", Status: "Pending"}
	time.Sleep(5 * time.Second)
	fmt.Println("All Done!!")
}