package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	i := 0
	log.Println(i)
	for i < 50 {
		log.Printf("logging from log-gen app. Index %d\n", i)
		log.Println("Sleep for 5 seconds")
		time.Sleep(5 * 1000 * time.Millisecond)
		i += 1
	}
}
