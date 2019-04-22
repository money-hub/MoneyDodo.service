package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
	http.ListenAndServe(":8080", nil)
}
