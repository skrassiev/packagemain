package b

import a "github.com/skrassiev/packagea"
import "fmt"

var B = 2

func init() {
	fmt.Println("Init package vendor/b")
	fmt.Println("b: a.A ==", a.A)
	fmt.Println("b: a.Inc() ==", a.Inc())
}
