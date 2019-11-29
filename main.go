package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	serveMineSweeper(port)
}
