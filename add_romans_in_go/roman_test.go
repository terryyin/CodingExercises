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

func ExampleXsAddI() {
	fmt.Println(romanAdd("XXXIX", "I"))
	fmt.Println(romanAdd("XLIX", "I"))
	fmt.Println(romanAdd("LXXXIX", "I"))
	fmt.Println(romanAdd("XCIX", "I"))
	// Output:
	// XL
	// L
	// XC
	// C
}

func ExampleCsAddI() {
	fmt.Println(romanAdd("CCCXCIX", "I"))
	fmt.Println(romanAdd("CMXCIX", "I"))
	// Output:
	// CD
	// M
}

func ExampleAddII() {
	fmt.Println(romanAdd("I", "II"))
	fmt.Println(romanAdd("II", "II"))
	// Output:
	// III
	// IV
}

func ExampleAddIII() {
	fmt.Println(romanAdd("I", "III"))
	// Output:
	// IV
}

func ExampleAddIV() {
	fmt.Println(romanAdd("I", "IV"))
	// Output:
	// V
}

func ExampleAddV() {
	fmt.Println(romanAdd("I", "V"))
	// Output:
	// VI
}

func ExampleAddVI() {
	fmt.Println(romanAdd("I", "VI"))
	// Output:
	// VII
}
