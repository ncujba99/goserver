package main

import (
	"fmt"
	"goserver/rest"
	"time"
)

func main() {
	go rest.Listen()

	for {
		fmt.Print("+")
		time.Sleep(10 * time.Second)
	}

}
