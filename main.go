package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/nearlynithin/monkey/repl"
)



func main() {
	user,err := user.Current()
	if err!= nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey console.\n",user.Username)
	fmt.Printf("Feel free to type commmands :)\n")
	repl.Start(os.Stdin,os.Stdout)
}