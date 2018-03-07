package intercom

import (
	"math"

	geo "github.com/kellydunn/golang-geo"
)

// Location is a particular place.
type Location interface {
	GeoPoint() *geo.Point
}

// Office is a location, usually a building or
// portion of a building, where a company conducts its business.
type Office struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// NewOffice returns a new Office given a name, latitude, and longitude.
func NewOffice(name string, lat, lng float64) Office {
	return Office{Name: name, Latitude: lat, Longitude: lng}
}

// GeoPoint returns a container object that holds
// geographic point coordinates.
func (o *Office) GeoPoint() *geo.Point {
	return geo.NewPoint(o.Latitude, o.Longitude)
}

// DistanceBetween finds distance between two Locations.
func DistanceBetween(l1, l2 Location) float64 {
	p1 := l1.GeoPoint()
	p2 := l2.GeoPoint()
	return math.Round(p1.GreatCircleDistance(p2))
}
