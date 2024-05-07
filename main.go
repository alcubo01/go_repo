// Ejercicio de inicialización de paquetes
package main

import (
	"fmt"
	"github.com/alcubo01/go_repo/mod1"
	"github.com/alcubo01/go_repo/mod2"
)

func init() {
	fmt.Println("Ready to launch ...")
}

func main() {
	saludo := "Liftoff! ACV"
	fmt.Println(saludo)
	fmt.Println(mod2.Translate("Arecibo message sent"))
	fmt.Println(mod1.Translate())
}
