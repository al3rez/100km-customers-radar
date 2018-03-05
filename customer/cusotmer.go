package customer

import (
	"strconv"

	geo "github.com/kellydunn/golang-geo"
)

type Customer struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	UserID    int    `json:"user_id"`
}

func NewCustomer() Customer {
	return Customer{}
}

func NewCustomerSlice() []Customer {
	return []Customer{}
}

func (c *Customer) CalcPoint() *geo.Point {
	lat, _ := strconv.ParseFloat(c.Latitude, 64)
	lng, _ := strconv.ParseFloat(c.Longitude, 64)
	return geo.NewPoint(lat, lng)
}
