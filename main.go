package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ishanz23/go-turso-starter-api/db"
)

func main() {
	start := time.Now()
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Time taken: ", time.Since(start))
	Server()
}
