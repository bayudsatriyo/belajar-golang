package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	reader, err := os.Open("Customer.json")

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	customer := &Customer{}

	decoder.Decode(customer)

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
}
