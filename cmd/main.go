package main

import (
	"fmt"

	"github.com/Masonchiiik/CalcGo/internal/application"
)

func main() {
	fmt.Printf("%v", "Hello user, have a nice day), and good luck in using the program (if something is unclear, read the README, it was made especially for you)\nThe main thing is not to close this console, otherwise the server will crash\n Here is a short guide for you if you are too lazy to read the readme: \n Just send your POST requests to this address: localhost:8080/api/v1/calculate\n")
	app := application.New()
	app.StartServer()
}
