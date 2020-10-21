package main

import (
	"gothic/study_003/myapp"
	"net/http"
)



func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
