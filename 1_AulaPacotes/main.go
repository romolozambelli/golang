package main

import (
	"fmt"
	"modulo/writeshell"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	writeshell.Writer()
	erro := checkmail.ValidateFormat("test@example.com")
	fmt.Println(erro)
}
