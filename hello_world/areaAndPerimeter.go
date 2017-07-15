package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.r * circle.r
}

func (rectangle Rectangle) area() float64 {
	length, width := rectangle.size()
	return length * width
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.r
}

func (rectangle Rectangle) perimeter() float64 {
	length, width := rectangle.size()
	return 2 * (length + width)

}
func (rectangle Rectangle) size() (float64, float64) {
	return math.Abs(rectangle.x2 - rectangle.x1),
		math.Abs(rectangle.y2 - rectangle.y1)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.perimeter()
	}
	return area
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{1, 1, 5, 5}
	fmt.Println(circle.area())
	fmt.Println(rectangle.area())
	fmt.Println(totalArea(circle, rectangle))
	fmt.Println(circle.perimeter())
	fmt.Println(rectangle.perimeter())
	fmt.Println(totalPerimeter(circle, rectangle))

}
