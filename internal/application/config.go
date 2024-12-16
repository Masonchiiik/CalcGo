package application

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Method avaible", 418)
		return
	}

	var requestBody Request
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, "Something wrong", 500)
		return
	}

	defer r.Body.Close()

	result, err := rpn.Calc(strings.TrimSpace(requestBody.Expression))
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	var ans Answer
	ans.Result = result
	jsonByte, err := json.Marshal(ans)

	if err != nil {
		http.Error(w, "Something wrong", 500)
		return
	}

	fmt.Fprint(w, string(jsonByte))
}

func (a *Application) StartServer() {
	mux := http.NewServeMux()
	Handler := http.HandlerFunc(CalculateHandler)

	mux.Handle("/api/v1/calculate", Handler)

	http.ListenAndServe(":8080", mux)
}
