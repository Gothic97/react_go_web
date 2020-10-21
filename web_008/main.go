package main

import (
	_ "fmt"
	"gothic/web_008/rest"
	"log"
)

func main() {

	log.Println("Start Server....")
	rest.RunAPI(":9090")
	log.Println("Finish Server....")
}
