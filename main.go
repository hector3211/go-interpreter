package main

import (
	"fmt"
	"go-interpreter/internal/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to the golang intrepreter %s\n", user.Username)
	fmt.Printf("Start typing in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
