package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {

	c1 := dog{"Rex", "Beagle", 6}
	c2 := dog{"Toby", "Dalmata", 5}

	fmt.Printf("The dog name is %s", c1.Name)
	fmt.Printf("The dog name is %s", c2.Name)

	dogC1InJSON, erro := json.Marshal(c1)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(dogC1InJSON)
	fmt.Println(bytes.NewBuffer(dogC1InJSON))

	dogC2InJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(dogC2InJSON)
	fmt.Println(bytes.NewBuffer(dogC2InJSON))

}
