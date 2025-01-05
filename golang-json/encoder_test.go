package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")	
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Eko",
		LastName: "Kurniawan",
		Age: 30,
		Married: true,
	}

	encoder.Encode(customer)

	writer.Close()

	fmt.Println(customer)
}
