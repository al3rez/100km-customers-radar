package json

import (
	"encoding/json"
	"io"
	"log"

	"github.com/azbshiri/100km-customers-radar/intercom"
)

// A CustomerReader reads and decodes JSON values
// from an input stream.
type CustomerDecoder struct {
	r io.Reader
}

// NewCustomerDecoder returns a new decoder that reads from r.
func NewCustomerDecoder(r io.Reader) *CustomerDecoder {
	return &CustomerDecoder{r: r}
}

// Decode reads the next JSON-encoded value from its
// input and stores it in the value pointed to by customers.
func (dec *CustomerDecoder) Decode(customers *[]intercom.Customer) {
	decoder := json.NewDecoder(dec.r)
	for {
		var customer intercom.Customer
		if err := decoder.Decode(&customer); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		*customers = append(*customers, customer)
	}
}
