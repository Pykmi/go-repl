package main

import (
	"fmt"
	"github.com/pykmi/repl/repl"
)

func main() {
	user := repl.Command{
		Name: "user",
		Action: func() error {
			fmt.Printf("User is %v", "pykmi")
			return nil
		},
		SubCommands: []repl.Command{
			{
				Name: "info",
				Action: func() error {
					fmt.Printf("User's name is %v", "pykmi")
					return nil
				},
			},
		},
	}

	replServer := repl.New()
	replServer.Register(user)

	replServer.Start()
}