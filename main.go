package main

import (
	"fmt"

	"github.com/a-poor/taks/lib"
)

func main() {
	t := lib.NewTask("Buy milk")
	fmt.Println(t)

	id, body, err := t.MarshalBytes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %s | Body: %s\n", string(id), string(body))
	fmt.Println("Is complete:", t.IsComplete())
}
