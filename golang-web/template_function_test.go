package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + myPage.Name
} 

func TemplateFunction(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Bayu"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Wowo",
	})
}

func TestTemplateFunction(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Wowo",
	})
}

func TestTemplateFunctionGlobal(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobalCustom(writer http.ResponseWriter, request *http.Request)  {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func (value string) string  {
			return strings.ToUpper(value)
		},
		"SayHello": func (value string) string  {
			return "Hello " + value
		},
	})

	t = template.Must(t.Parse(`{{ SayHello .Name | upper }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Wowo Ngeri",
	})
}

func TestTemplateFunctionGlobalCustom(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalCustom(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}