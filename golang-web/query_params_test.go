package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")

	if query == "" {
		fmt.Fprintln(w, "Hello World")
	} else {
		fmt.Fprintf(w, "Hello %s", query)
	}
}

func TestQueryParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=ade", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
} 

func MultipleQueryParams(writer http.ResponseWriter, request *http.Request) {
	firstname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestMultipleQueryParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?first_name=ade&last_name=kheneddy", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParams(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParamsValue(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	names := query["name"]

	fmt.Println(writer, strings.Join(names, " "))
}

func TestMultipleParamsValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=ade&name=kheneddy", nil)
	recorder := httptest.NewRecorder()

	MultipleParamsValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}