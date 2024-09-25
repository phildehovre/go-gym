package main

import "log"

func main() {
	server := NewAPIServer(":5000")

	err := server.Run()

	if err != nil {
		log.Fatalf("error running server %v", err)
	}
}
