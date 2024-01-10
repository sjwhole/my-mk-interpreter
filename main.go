package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 2 {
		fileName := os.Args[1]
		open, err := os.Open(fileName)
		if err != nil {
			return
		}
		repl.StartWithFile(open, os.Stdout)
	} else {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
