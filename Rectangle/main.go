package main

import (
	"errors"
	"fmt"
)

//Create a rectangle abstraction using struct.
//Create two functions to calculate area and circumference of given rectangle instance and set produced values on given rectangle instance.
//Create another function to create an instance of rectangle struct and return it. But that function would be able to return an error in case of passing invalid arguments such as negative length or height.
//Choose one of strategies mentioned i the text book and exemplified in error.go of ch 05 in example project.

// Rectangle structure
type Rectangle struct {
	height int
	length int
}

//Function to make a proper Rectangle
func makeRec(length, height int) (Rectangle, error) {

	rectangle1 := Rectangle{
		height: height,
		length: length,
	}

	if height <= 0 || length <= 0 {
		err := errors.New("Negative values are not available!")
		return rectangle1, err
	}

	return rectangle1, nil

}

//Area of Rectangle
func Area(rectangle Rectangle) int {
	return rectangle.length * rectangle.height
}

//Circumference of Rectangle
func Circumference(rectangle Rectangle) int {
	return 2 * (rectangle.length + rectangle.height)
}

func main() {
	//Blank identifier to ignore error
	rectangle, _ := makeRec(10, 8)

	fmt.Println(Area(rectangle))

	fmt.Println(Circumference(rectangle))
}
