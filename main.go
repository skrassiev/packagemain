package main

import (
	"fmt"
	"github.com/skrassiev/packagea"
	"github.com/skrassiev/packageb"
	"github.com/skrassiev/packagec"
	"io/ioutil"
)

func main() {
	fmt.Println("main: packagea.A ==", packagea.A)
	fmt.Println("main.main()")
	fmt.Println("main: package.Inc() ==", packagea.Inc())
	fmt.Println("main: packageb.B ==", packageb.B)
	fmt.Println("main: packagec.C ==", packagec.Caap)
	packagec.IncC()
	fmt.Println(ioutil.Discard)
}
