package main

import (
	"net/http"
)

type HelloHandle struct{}

func (h *HelloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.Handle("/", &HelloHandle{})
	//_ = http.ListenAndServe(":12345", nil)

	//(2)More control over the server's behavior is available by creating a custom Server:
	//
	s := &http.Server{
		Addr: ":12345",
	}
	_ = s.ListenAndServe()
}
