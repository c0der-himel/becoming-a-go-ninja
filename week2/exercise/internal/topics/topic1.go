// Example: A ride-sharing app (like Uber)

package topics

import "fmt"

// Struct = a data model
type Car struct {
	Brand string
	Seats int
}

// Interface = defines behavior
type Vehicle interface {
	Drive() string
}

// Implement the interface for Car
func (c Car) Drive() string {
	return fmt.Sprintf("%s is driving with %d seats", c.Brand, c.Seats)
}

type Bike struct {
	Brand string
}

func (b Bike) Drive() string {
	return fmt.Sprintf("%s bike is riding . . .", b.Brand)
}

func Vehicles() {
	var v Vehicle

	v = Car{Brand: "BMW", Seats: 5}
	fmt.Println(v.Drive())

	v = Bike{Brand: "Honda"}
	fmt.Println(v.Drive())
}
