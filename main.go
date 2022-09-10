package main

import (
	"com/tutorials/helpers"
	"fmt"
)

func main() {
	fmt.Println("First go program")

	var myStudnet helpers.Student

	myStudnet.Name = " name sirname"
	myStudnet.RollNumber = 45
	fmt.Printf("Stuend name is %s and his roll number is %d", myStudnet.Name, myStudnet.RollNumber)
}
