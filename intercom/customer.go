package intercom

import (
	"encoding/json"
	"strconv"

	geo "github.com/kellydunn/golang-geo"
)

// A Customer is a person of a specified kind
// that Office has to deal with.
type Customer struct {
	UserID    int     `json:"user_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GeoPoint returns a container object that
// holds geographic point coordinates.
func (c *Customer) GeoPoint() *geo.Point {
	return geo.NewPoint(c.Latitude, c.Longitude)
}

// UnmarshalJSON sets *c to a copy of data to deserialize custom types.
func (c *Customer) UnmarshalJSON(data []byte) error {
	// aliasedCustomer is a alias of Customer in which have all the
	// same fields, but none of the methods to avoid inherit the
	// original UnmarshalJSON method, causing it to go into an infinite loop.
	type aliasedCustomer Customer

	aux := &struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
		*aliasedCustomer
	}{
		aliasedCustomer: (*aliasedCustomer)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	c.Latitude, _ = strconv.ParseFloat(aux.Latitude, 64)
	c.Longitude, _ = strconv.ParseFloat(aux.Longitude, 64)

	return nil
}
