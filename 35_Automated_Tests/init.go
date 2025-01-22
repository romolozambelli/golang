package main

import (
	"fmt"

	"init.go/endereco"
)

func main() {

	tipoEndereco := endereco.AddressTypeCheck("Road Stewart")

	fmt.Println("Hello World")
	fmt.Println(tipoEndereco)

}
