package main

import (
	"fmt"

	"github.com/Masonchiiik/CalcGo/internal/application"
)

func main() {
	fmt.Printf("%v", "Привет пользователь, хорошего тебе дня), и удачи в пользовании программы(если что-то непонятно почитай README, он сделан специально для тебя)\nГлавное не закрывай эту консольку, а то сервер упадёт\n Вот тебе краткий гайд если лень читать ридми:\n Просто отправляй свои POST запросы на этот адрес: localhost/api/v1/calculate\n")
	app := application.New()
	app.StartServer()
}
