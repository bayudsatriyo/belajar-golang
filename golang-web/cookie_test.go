package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(write http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Powered-By"
	cookie.Value = "Bayu"
	cookie.Path = "/"

	http.SetCookie(write, cookie)
	fmt.Fprint(write, "Success Create Cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Powered-By")

	if err != nil {
		fmt.Fprintln(writer, "Cookie Not Found")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=eko", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	response := recorder.Result()
	cookie := response.Cookies()

	for _, cookies := range cookie {
		fmt.Println(cookies.Name)
		fmt.Println(cookies.Value)
	}
}

func TestGetCookie(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookies := new(http.Cookie)
	cookies.Name = "X-Powered-By"
	cookies.Value = "Bayu"
	request.AddCookie(cookies)
	
	response := httptest.NewRecorder()

	GetCookie(response, request)

	body, _ := io.ReadAll(response.Result().Body)

	fmt.Println(string(body))
}