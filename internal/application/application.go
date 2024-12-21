package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Masonchiiik/CalcGo/pkg/rpn"
)

type Application struct {
}

type Answer struct {
	Result float64 `json:"result"`
}

type Request struct {
	Expression string `json:"expression"`
}

func New() *Application {
	return &Application{}
}

func CheckMethodMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, errorMethod, http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {

	file, err := os.OpenFile("CalcGo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	if err != nil {
		log.Print("Server end work with 500 code")
		http.Error(w, errorInternal, 500)
	}

	var requestBody Request
	err = json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		log.Print("Server end work with 500 code")
		http.Error(w, errorInternal, 500)
		return
	}

	defer r.Body.Close()

	result, err := rpn.Calc(strings.TrimSpace(requestBody.Expression))
	if err != nil {
		log.Print("Server end work with 422 code")
		http.Error(w, errorExpression, 422)
		return
	}

	var res Answer
	res.Result = result
	jsonByte, err := json.Marshal(res)

	if err != nil {
		log.Print("Server end work with 500 code")
		http.Error(w, errorInternal, 500)
		return
	}

	fmt.Fprint(w, string(jsonByte))
	log.Printf("server end work with code 400 and with result: %v", res.Result)
}

func (a *Application) StartServer() {
	mux := http.NewServeMux()
	Handler := http.HandlerFunc(CalculateHandler)

	mux.Handle("/api/v1/calculate", CheckMethodMiddlerware(Handler))

	http.ListenAndServe("", mux)
}
