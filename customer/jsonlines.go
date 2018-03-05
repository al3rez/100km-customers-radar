package customer

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

func Unmarshal(data []byte, customers *[]Customer) {
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	for {
		var customer Customer
		if err := decoder.Decode(&customer); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		*customers = append(*customers, customer)
	}
}
