package golangweb

import (
	"net/http"
	"testing"
)


func TestService(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}