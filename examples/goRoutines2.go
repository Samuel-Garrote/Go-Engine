/*
	 func main(){
		id:=3
		result := fmt.Sprintf("Job %d done", id)
		fmt.Println(result)
	}
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func processJob(id int, wg *sync.WaitGroup, ch chan string) {
    defer wg.Done()
    time.Sleep(1 * time.Second)
    ch <- fmt.Sprintf("Job %d done", id)
}

func main7() {
    var wg sync.WaitGroup
    ch := make(chan string, 5)

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go processJob(i, &wg, ch)
    }

    wg.Wait()
    close(ch)

    for result := range ch {
        fmt.Println(result)
    }
}