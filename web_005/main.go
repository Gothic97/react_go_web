package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("src/gothic/web_005/public")))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
