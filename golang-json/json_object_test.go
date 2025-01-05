package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName string
	Age int
	Married bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Eko",
		LastName: "Kurniawan",
		Age: 30,
		Married: true,
	}

	bytes, err := json.Marshal(customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
