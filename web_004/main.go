package main

import (
	"gothic/web_004/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())

}
