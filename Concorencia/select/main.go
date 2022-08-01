package main

import (
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	go func() {

		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

}
