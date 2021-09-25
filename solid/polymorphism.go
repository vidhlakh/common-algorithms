package solid

// defining an interface
type Figure interface {
	Area() float64
}

// declaring a struct
type Rectangle struct {

	// declaring struct variables
	Length float64
	Width  float64
}

// declaring a struct
type Square struct {

	// declaring struct variable

	Side float64
}

// function to calculate
// area of a rectangle
func (rect Rectangle) Area() float64 {

	// Area of rectangle = l * b
	area := rect.Length * rect.Width
	return area
}

// function to calculate
// area of a square
func (sq Square) Area() float64 {

	// Area of square = a * a
	area := sq.Side * sq.Side
	return area
}

// // main function
// func main() {

// 	// declaring a rectangle instance
// 	rectangle := Rectangle{

// 		length: 10.5,
// 		width:  12.25,
// 	}

// 	// declaring a square instance
// 	square := Square{

// 		side: 15.0,
// 	}

// 	// printing the calculated result
// 	fmt.Printf("Area of rectangle: %.3f unit sq.\n", rectangle.Area())
// 	fmt.Printf("Area of square: %.3f unit sq.\n", square.Area())
// }
