package main

import (
	"errors"
	"fmt"
)

//Create a rectangle abstraction using struct.
//Create two functions to calculate area and circumference of given rectangle instance and set produced values on given rectangle instance.
//Create another function to create an instance of rectangle struct and return it. But that function would be able to return an error in case of passing invalid arguments such as negative length or height.
//Choose one of strategies mentioned i the text book and exemplified in error.go of ch 05 in example project.

//Struct of rectangle
type Rectangle struct {
	height int
	len    int
}

//Function to make a proper Rectangle
func makeRec(len, heig int) (Rectangle, error) {

	rect1 := Rectangle{
		height: heig,
		len:    len,
	}

	if heig <= 0 || len <= 0 {
		err := errors.New("Negative values are not available!")
		return rect1, err
	}

	return rect1, nil

}

//Area of Rectangle
func Area(rectangle Rectangle) int {
	return rectangle.len * rectangle.height
}

//Circumference of Rectangle
func Circumference(rectangle Rectangle) int {
	return 2 * (rectangle.len + rectangle.height)
}

func main() {
	//Blank identifier to ignore err
	rect1, _ := makeRec(10, 8)

	//Area calculation
	fmt.Println(Area(rect1))

	//Circumference
	fmt.Println(Circumference(rect1))
}
