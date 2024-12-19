package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, errorPost, 405)
			return
		}

		var body Request

		file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			http.Error(w, errorInternal, 500)
			return
		}
		defer file.Close()

		log.SetOutput(file)

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, errorInternal, 500)
			return
		}

		r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

		err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&body)
		if err != nil {
			http.Error(w, errorInternal, 500)
			return
		}

		log.Printf("Body: %v", body)

		next.ServeHTTP(w, r)
	})
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {

	var requestBody Request
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, errorInternal, 500)
		return
	}

	defer r.Body.Close()

	result, err := rpn.Calc(strings.TrimSpace(requestBody.Expression))
	if err != nil {
		http.Error(w, errorExpression, 422)
		return
	}
	var res Answer
	res.Result = result
	jsonByte, err := json.Marshal(res)

	if err != nil {
		http.Error(w, errorInternal, 500)
		return
	}

	fmt.Fprint(w, string(jsonByte))
}

func (a *Application) StartServer() {
	mux := http.NewServeMux()
	Handler := http.HandlerFunc(CalculateHandler)

	mux.Handle("/api/v1/calculate", LoggingMiddleware(Handler))

	http.ListenAndServe("", mux)
}
