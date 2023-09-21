package main

import (
	"fmt"
	"homework/shapes"
)

func main() {
	circle := shapes.Circle{Radius: 5}
	rectangle := shapes.Rectangle{Width: 4, Height: 3}

	composite := shapes.CompositeShape{Shapes: []shapes.Shape{circle, rectangle}}

	fmt.Printf("Площадь круга: %f\n", circle.Area())
	fmt.Printf("Периметр круга: %f\n", circle.Perimeter())

	fmt.Printf("Площадь прямоугольника: %f\n", rectangle.Area())
	fmt.Printf("Периметр прямоугольника: %f\n", rectangle.Perimeter())

	fmt.Printf("Площадь составной фигуры: %f\n", composite.Area())
	fmt.Printf("Периметр составной фигуры: %f\n", composite.Perimeter())
}
