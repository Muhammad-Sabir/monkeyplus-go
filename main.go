package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Muhammad-Sabir/monkeyplus-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s! Welcome to Monkey Plus language! \n", user.Username)
	fmt.Printf("Lets get started with the commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
