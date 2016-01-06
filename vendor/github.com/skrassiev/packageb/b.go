package packageb

import a "github.com/skrassiev/packagea"
import "fmt"

var B = 2

func IncB() int {
	B++
	return B
}

func init() {
	fmt.Println("Init package vendor/b")
	fmt.Println("b: a.A ==", a.A)
	fmt.Println("b: a.Inc() ==", a.Inc())
}
