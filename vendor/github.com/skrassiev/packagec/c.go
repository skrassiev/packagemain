package c

import a "github.com/skrassiev/packagea"
import "fmt"

var C = 3

func init() {
	fmt.Println("Init package vendor/c")
	fmt.Println("c: C ==", C)
	fmt.Println("c: a.A ==", a.A)
	fmt.Println("c: a.Inc() ==", a.Inc())
}
