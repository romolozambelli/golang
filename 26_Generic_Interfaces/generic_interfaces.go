package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("home")
	generic(123)
	generic(true)

	mapa := map[interface{}]interface{}{
		1:            "casa",
		float32(100): true,
	}

	fmt.Println(mapa)
}
