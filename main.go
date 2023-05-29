package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for {
		s := "This is a memory leak example" // Allocating memory in an infinite loop
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}
