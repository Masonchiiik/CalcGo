package main

import (
	"fmt"

	"github.com/Masonchiiik/CalcGo/internal/application"
)

func main() {
	fmt.Printf("%v", "Привет лицеист, хорошего тебе дня), и удачи в пользовании программы(если что-то непонятно почитай README, он сделан специально для тебя)")
	app := application.New()
	app.StartServer()
}
