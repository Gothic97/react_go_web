package main

import (
	"log"

	"gothic/web_014/backend/src/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
