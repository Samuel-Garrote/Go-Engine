package main

import (
	"fmt"
)

func sum2(a int, b int, ch chan int ){
		ch <- a+b
}

func main5(){
	ch := make(chan int)

	go sum2(5,3, ch)
	result := <-ch
	fmt.Println(result)
}