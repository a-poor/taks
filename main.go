package main

import (
	"os"

	"github.com/a-poor/taks/cmd"
)

func main() {
	// t := lib.NewTask("Buy milk")
	// t.Priority = lib.TaskPriorityHigh
	// fmt.Println(t)

	// id, body, err := t.MarshalBytes()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("ID: %s | Body: %s\n", string(id), string(body))
	// fmt.Println("Is complete:", t.IsComplete())

	app := cmd.NewApp()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
