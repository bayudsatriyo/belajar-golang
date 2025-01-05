package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":1,"name":"Apple Macbook","price":1000000}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}

	err := json.Unmarshal(jsonByte, &result)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["name"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        1,
		"name":      "Apple Macbook",
		"price":     1000000,
		"image_url": "macbook.com",
	}

	bytes, err := json.Marshal(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}