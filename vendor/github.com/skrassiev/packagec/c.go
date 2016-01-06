package packagec

import a "github.com/skrassiev/packagea"
import "fmt"

var C = 3
var Caap = 3

func IncC() int {
	C++
	return C
}

func init() {
	fmt.Println("Init package vendor/c")
	fmt.Println("c: C ==", C)
	fmt.Println("c: a.A ==", a.A)
	fmt.Println("c: a.Inc() ==", a.Inc())
}
