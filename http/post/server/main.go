package main

import "net/http"

func main() {

	_ = http.ListenAndServe(":12345", nil)
}
