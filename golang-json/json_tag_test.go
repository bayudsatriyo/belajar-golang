package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id: 1,
		Name: "Eko",
		ImageUrl: "eko.com",
	}

	bytes, err := json.Marshal(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonDecode(t *testing.T) {
	jsonString := `{"id":1,"name":"Eko","image_url":"eko.com"}`
	jsonByte := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonByte, product)

	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Name)
}