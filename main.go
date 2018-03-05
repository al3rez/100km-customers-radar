package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"

	"github.com/azbshiri/invite-customers-within-100km/customer"
	geo "github.com/kellydunn/golang-geo"
)

func main() {
	allCustomers := customer.NewCustomerSlice()
	data, _ := ioutil.ReadFile("customers.jsonl")
	customer.Unmarshal(data, &allCustomers)

	dublinOffice := geo.NewPoint(53.339428, -6.257664)
	customersWithin100KM := customer.NewCustomerSlice()
	for _, customer := range allCustomers {
		distance := distanceBetween(dublinOffice, customer.CalcPoint())
		if (distance >= 0) && (distance <= 100) {
			customersWithin100KM = append(customersWithin100KM, customer)
		}
	}

	// Sort customersWithin100KM of dublinOffice by UserID (ascending)
	sort.Slice(customersWithin100KM[:], func(i, j int) bool {
		return customersWithin100KM[i].UserID < customersWithin100KM[j].UserID
	})

	// Output the Names and UserIDs of customersWithin100KM
	for _, customer := range customersWithin100KM {
		fmt.Printf("%d\t %s\n", customer.UserID, customer.Name)
	}
}

func distanceBetween(p1, p2 *geo.Point) float64 {
	return math.Round(p1.GreatCircleDistance(p2))
}
