package main

import "fmt"

func main() {

	//In order words. Variables are just place holders
	// We can setup variables many ways in go

	var favoriteFood string = "Pizza"
	name := "Kendy" //This will store the value
	age := 24

	// also := this will assign the value type to the variable

	fmt.Println("My name is ", name)
	fmt.Println("My age is ", age)

	fmt.Printf("My name is %v, i'm %d years old and my favorite food is %v", name, age, favoriteFood)

}
