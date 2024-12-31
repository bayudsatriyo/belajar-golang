package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
    Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
    fmt.Println("Before execute handler")
    middleware.Handler.ServeHTTP(writer, request)
    fmt.Println("After execute handler") 
}

type ErrorMiddleware struct {
    Hanlder http.Handler
}

func (handler *ErrorMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
    defer func() {
        err := recover()
        fmt.Println(err)
        if err != nil {
            writer.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(writer, "ERROR : %s", err)
        }
    }()

    handler.Hanlder.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T)  {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintln(writer, "Hello World")
        fmt.Fprintln(writer, "Hello Middleware")
    })

    mux.HandleFunc("/panic", func (writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintln(writer, "Hello Panic")
        panic("Ups Error")
    })

     logMiddleware := &LogMiddleware{
        Handler: mux,
    }

    errorMiddleware := &ErrorMiddleware{
        Hanlder: logMiddleware,
    }

    server := http.Server{
        Addr: "localhost:8080",
        Handler: errorMiddleware,
    }

    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}