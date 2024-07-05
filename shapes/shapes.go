package shapes

import "math"

type Shape interface {
    Area() float64
    Double()
}

type Square struct {
    sideLength float64
}

func (s *Square) Area() float64 {
    return s.sideLength * s.sideLength
}

func (s *Square) Double() {
    s.sideLength *= 2
}

type Circle struct {
    radius float64
}

func (c *Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c *Circle) Double() {
    c.radius *= 2
}
