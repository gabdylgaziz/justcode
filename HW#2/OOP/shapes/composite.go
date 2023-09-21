package shapes

type CompositeShape struct {
	Shapes []Shape
}

func (cs CompositeShape) Area() float64 {
	totalArea := 0.0
	for _, shape := range cs.Shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func (cs CompositeShape) Perimeter() float64 {
	totalPerimeter := 0.0
	for _, shape := range cs.Shapes {
		totalPerimeter += shape.Perimeter()
	}
	return totalPerimeter
}
