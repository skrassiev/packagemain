package packagea

import "fmt"

var A = 1

func Inc() int {
	A++
	return A
}

func init() {
	fmt.Println("Init package vendor/a")
	fmt.Println("a: A ==", A)
}
