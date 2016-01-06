package main

import (
	"fmt"
	a "github.com/skrassiev/packagea"
	b "github.com/skrassiev/packageb"
	c "github.com/skrassiev/packagec"
)

func main() {
	fmt.Println("main.main()")
	fmt.Println("main: a.A ==", a.A)
	fmt.Println("main: a.Inc() ==", a.Inc())
	fmt.Println("main: b.B ==", b.B)
	fmt.Println("main: c.C ==", c.C)
}
