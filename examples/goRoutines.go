package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Starting:", name)
    time.Sleep(2 * time.Second)
    fmt.Println("Done:", name)
}

func main3() {
    var wg sync.WaitGroup

    wg.Add(3)
    go task("Task 1", &wg)
    go task("Task 2", &wg)
    go task("Task 3", &wg)

    wg.Wait()
    fmt.Println("All tasks done!")
}