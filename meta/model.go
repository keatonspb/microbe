package meta

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

type Shape struct {
	Start Point
	End   Point
}

type Vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func NewShape(oX float64, oY float64, oH float64, oW float64) Shape {
	return Shape{
		Start: Point{X: oX, Y: oY},
		End:   Point{X: oX + oW, Y: oY + oH},
	}
}
