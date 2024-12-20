package application_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Masonchiiik/CalcGo/internal/application"
)

type RequestTest struct {
	Expression string `json:"expression"`
}
type ResponseTest struct {
	Result float64 `json:"result"`
}

func TestCalculateHandler(t *testing.T) {

	tests := []struct {
		name           string
		input          RequestTest
		expectedResult float64
		err            error
	}{
		{
			name: "Valid expression",
			input: RequestTest{
				Expression: "2+2",
			},
			expectedResult: 4,
			err:            nil,
		},
		{
			name: "Invalid expression",
			input: RequestTest{
				Expression: "2+",
			},
			err: nil,
		},
		{
			name: "Invalid expression",
			input: RequestTest{
				Expression: "2+()",
			},
			err: errors.New("Expression is not valid"),
		},
		{
			name: "Valid expression",
			input: RequestTest{
				Expression: "2+2*2",
			},
			expectedResult: 6,
			err:            nil,
		},
		{
			name: "Invalid expression",
			input: RequestTest{
				Expression: "()",
			},
			err: errors.New("Expression is not valid"),
		},
		{
			name: "Valid expression",
			input: RequestTest{
				Expression: "1024*2",
			},
			expectedResult: 2048,
			err:            nil,
		},
		{
			name: "Valid expression",
			input: RequestTest{
				Expression: "1/2",
			},
			expectedResult: 0.5,
			err:            nil,
		},
		{
			name: "Valid expression with float input and with space",
			input: RequestTest{
				Expression: "0.5 + 0.5",
			},
			expectedResult: 1,
			err:            nil,
		},
		{
			name: "Valid expression with float input and with space",
			input: RequestTest{
				Expression: "3.142134132 / 2",
			},
			expectedResult: 1.571067066,
			err:            nil,
		},
		{
			name: "Valid expression",
			input: RequestTest{
				Expression: "1+2",
			},
			expectedResult: 3,
			err:            nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(application.CalculateHandler)
			body, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatal("json failed")
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			var resp ResponseTest
			err = json.NewDecoder(rec.Body).Decode(&resp)
			if err != nil {
				if err != tt.err {
					t.Errorf("error is not valid")
				}
			}

			if resp.Result != tt.expectedResult {
				t.Errorf("expected result %v, got %v", tt.expectedResult, resp.Result)
			}
		})
	}
}
