package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("assert/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8081", nil)
}
