package main

import( 
	"fmt"
	"math"
)

// Implement types Rectangle, Circle and Shape
type Shape interface {
	Area() float64
}

type Square struct {
	width  float64
	length float64
}

func (s Square) Area() float64 {
	return s.width * s.length
}

func (s Square) String() string {
	return fmt.Sprintf("Square of width: %.2f and length: %.2f", s.width, s.length)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle of Radius: %.2f", c.radius)
}

func main() {
	r := &Square{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
