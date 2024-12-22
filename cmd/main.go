package main

import (
	"fmt"

	"github.com/Masonchiiik/CalcGo/internal/application"
)

func main() {
	fmt.Printf("%s", "Server is starting...\nWarning: Do not close this console!\nIf don't understand how to work with server, please, read README.md\n")
	app := application.New()
	app.StartServer()
}
