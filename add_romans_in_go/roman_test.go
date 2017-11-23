package roman

import "fmt"

func ExampleIsAddI() {
	fmt.Println(romanAdd("I", "I"))
	fmt.Println(romanAdd("II", "I"))
	fmt.Println(romanAdd("III", "I"))
	fmt.Println(romanAdd("IV", "I"))
	fmt.Println(romanAdd("V", "I"))
	fmt.Println(romanAdd("VI", "I"))
	fmt.Println(romanAdd("VII", "I"))
	fmt.Println(romanAdd("VIII", "I"))
	fmt.Println(romanAdd("IX", "I"))
	// Output:
	// II
	// III
	// IV
	// V
	// VI
	// VII
	// VIII
	// IX
	// X
}

func ExampleXAddI() {
	fmt.Println(romanAdd("X", "I"))
	fmt.Println(romanAdd("XI", "I"))
	fmt.Println(romanAdd("XII", "I"))
	fmt.Println(romanAdd("XIII", "I"))
	fmt.Println(romanAdd("XIV", "I"))
	fmt.Println(romanAdd("XIX", "I"))
	fmt.Println(romanAdd("XX", "I"))
	// Output:
	// XI
	// XII
	// XIII
	// XIV
	// XV
	// XX
	// XXI
}
