package main
import "fmt"
type Shape interface {
    Area() float64
}

type Rectangle struct {
    width  float64
    height float64
}

type Square struct {
    side float64
}

func (r *Rectangle) Area() float64 {
    return r.width * r.height
}

func (s *Square) Area() float64 {
    return s.side * s.side
}

func PrintArea(shape Shape) {
    fmt.Println(shape.Area())
}

func main() {
    rectangle := &Rectangle{width: 5, height: 10}
    PrintArea(rectangle)

    square := &Square{side: 5}
    PrintArea(square)
}