package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)



type Action struct {
	Title string
	Name string
}

type Compare struct {
	Title string
	FinalScore int
}

type Hobbies struct {
	Title string
	Hobbie []string
}

func TemplateDataAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Action{
		Title: "Template Data Struct",
		Name: "",
	})
}

func TestTemplateDataAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataAction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateDataComaparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", Compare{
		Title: "Template Data Struct",
		FinalScore: 60,
	})
}

func TestTemplateDataComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataComaparator(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateDataRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", Hobbies{
		Title: "Template Data Struct",
		Hobbie: []string{"Coding", "Reading", "Sleeping"},
	})
}

func TestTemplateDataRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

type AddressDetail struct {
	Street string
	City string
}

type AddressWith struct {
	Title string
	Name string
	Address AddressDetail
}

func TemplateDataWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Name": "Bayu",
		"Title": "Template Data Map",
		
	})
}

func TestTemplateDataWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataWith(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}