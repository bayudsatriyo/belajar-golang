package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		http.ServeFile(w, r, "./resources/index.html")
	} else {
		fmt.Fprintf(w, "Empty file %s", name)
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources/index.html
var resourcesOk string

//go:embed resources/index.css
var resourcesNotOk string

func ServeFileGoEmbed(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name != "" {
		fmt.Fprint(w, resourcesOk)
	} else {
		fmt.Fprint(w, resourcesNotOk)
	}
}

func TestServeFileGoEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileGoEmbed),
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}