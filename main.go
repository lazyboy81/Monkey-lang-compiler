package main

import (
	"Monkey-Lang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	userInfo, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		userInfo.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
