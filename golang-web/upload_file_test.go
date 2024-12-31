package golangweb

import (
	"net/http"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request)  {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}