package main

import (
	"fmt"
	"monkeyplus-go/repl"
	"os"
	"os/user"
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
