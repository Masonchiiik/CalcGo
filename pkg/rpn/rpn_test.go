package rpn_test

import (
	"testing"

	"github.com/Masonchiiik/CalcGo/pkg/rpn"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "simple",
			expression:     "2+2+3",
			expectedResult: 7,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
		{
			name:           "priority",
			expression:     "2+2+(2*2)/2*(2+2)",
			expectedResult: 12.0,
		},
		{
			name:           "simple",
			expression:     "2+2",
			expectedResult: 4.0,
		},
		{
			name:           "one char",
			expression:     "2",
			expectedResult: 2.0,
		},
		{
			name:           "brackets",
			expression:     "(2*2)/4",
			expectedResult: 1.0,
		},
		{
			name:           "addition and division",
			expression:     "4+6/3",
			expectedResult: 6.0,
		},
		{
			name:           "single number",
			expression:     "42",
			expectedResult: 42.0,
		},
	}

	for _, test_case := range testCasesSuccess {
		t.Run(test_case.name, func(t *testing.T) {
			val, err := rpn.Calc(test_case.expression)

			if err != nil {
				t.Fatalf("succesful case %v error", test_case.expression)
			}

			if val != test_case.expectedResult {
				t.Fatalf("%f should be equal %f", val, test_case.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "simple",
			expression: "1+1*",
		},
		{
			name:       "priority",
			expression: "2+2**2",
		},
		{
			name:       "priority",
			expression: "((2+2-*(2",
		},
		{
			name:       "empty",
			expression: "",
		},
		{
			name:       "simple",
			expression: "+2-+2-2+23+-",
		},
		{
			name:       "division by zero",
			expression: "1/0",
		},
		{
			name:       "too many operators",
			expression: "2***3",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := rpn.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %v is invalid but result  %v", testCase.expression, val)
			}
		})
	}
}
