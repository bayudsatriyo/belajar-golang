package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncodeJson(t *testing.T) {
	logJson(map[string]string{
		"foo": "bar",
		"baz": "qux",
	})
	logJson("Eko")
	logJson(true)
	logJson(2)
	logJson([]string{"Eko", "Kurniawan", "Khannedy"})
}
