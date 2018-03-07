package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/azbshiri/100km-customers-radar/intercom"
	"github.com/azbshiri/100km-customers-radar/internal/json"
)

func main() {
	customers := []intercom.Customer{}
	r, _ := os.Open("customers.jsonl")
	decoder := json.NewCustomerDecoder(r)
	decoder.Decode(&customers)

	dublinOffice := intercom.NewOffice("Dublin", 53.339428, -6.257664)
	customersWithin100KM := []intercom.Customer{}
	for _, customer := range customers {
		distance := intercom.DistanceBetween(&dublinOffice, &customer)
		if distance <= 100 {
			customersWithin100KM = append(customersWithin100KM, customer)
		}
	}

	sort.Slice(customersWithin100KM[:], func(i, j int) bool {
		return customersWithin100KM[i].UserID < customersWithin100KM[j].UserID
	})

	for _, customer := range customersWithin100KM {
		fmt.Printf("%d\t %s\n", customer.UserID, customer.Name)
	}
}
