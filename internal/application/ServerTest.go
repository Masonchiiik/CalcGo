package application

import (
	"bytes"
	"encoding/json"
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
	handler := http.HandlerFunc(application.CalculateHandler)

	tests := []struct {
		name       string
		input      RequestTest
		expected   float64
		statusCode int
	}{
		{
			name: "Normal expression",
			input: RequestTest{
				Expression: "2+2",
			},
			expected:   4,
			statusCode: http.StatusOK,
		},
		{
			name: "Incomplete expression",
			input: RequestTest{
				Expression: "2+",
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid characters in expression",
			input: RequestTest{
				Expression: "2+()",
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal input: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if rec.Code != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, rec.Code)
			}

			if tt.statusCode == http.StatusOK {
				var resp ResponseTest
				err := json.NewDecoder(rec.Body).Decode(&resp)
				if err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				if resp.Result != tt.expected {
					t.Errorf("expected result %f, got %f", tt.expected, resp.Result)
				}
			}
		})
	}
}
