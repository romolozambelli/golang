package main

import (
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

	var c1 dog

	dogC1InJSON := `{"name":"Toby","breed":"Dalmata","age":5}`

	if erro := json.Unmarshal([]byte(dogC1InJSON), &c1); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c1)

	dogC2InJSON := `{"name":"Mike","breed":"Pitbull"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(dogC2InJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}
