package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length int
	width int
}

func (r *Rectangle) Area() float64 {
	return float64(r.length * r.width)
}

func (r *Rectangle) Perimeter() float64 {
	return float64(2 * r.length + 2 * r.width) 
}

type Circle struct {
	radius int
}

func (c *Circle) Area() float64 {
	return 3.14 * float64(c.radius * c.radius)
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * float64(c.radius)
}

func PrintShapeDetails(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

type Triangle struct {
    base   int
    height int
}

func (t *Triangle) Area() float64 {
    return 0.5 * float64(t.base) * float64(t.height)
}

func main() {
	fmt.Println("checking")
	r := &Rectangle{
		length: 2,
		width: 4,
	}

	c := &Circle{
		radius: 4,
	}

	t := &Triangle{
		base: 2,
		height: 3,
	}

	// fmt.Println("Printing Circle area ", c.Area(), " Perimeter ", c.Perimeter())
	// fmt.Println("Printing Rectangle area ", r.Area(), " Perimeter ", r.Perimeter())

	fmt.Println(t.Area())
	PrintShapeDetails(r)
	PrintShapeDetails(c)
}