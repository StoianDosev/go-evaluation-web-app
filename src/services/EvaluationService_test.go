package services

import (
	"fmt"
	"testing"
)

type TestData struct {
	input        string
	expected     int
	errorMessage string
	valid        bool
}

func TestCalculateExpressionWithCorrectInput(t *testing.T) {
	// Arrange
	dataInput := []TestData{
		{input: "What is 5 plus 13?", expected: 18},
		{input: "What is 7 minus 5?", expected: 2},
		{input: "What is 6 multiplied by 4?", expected: 24},
		{input: "What is 25 divided by 5?", expected: 5},
		{input: "What is 25?", expected: 25},
		{input: "What is 3 plus 2 multiplied by 3?", expected: 15},
	}

	sut := NewEvaluationService()

	for _, v := range dataInput {

		// Act
		result, err := sut.CalculateExpression(v.input)

		// Assert
		if result != v.expected || err != nil {
			t.Errorf("Not correct value of: %v, expexted %v", result, v.expected)
		}
	}
}

func TestCalculateExpressionWithUnsupportedInput(t *testing.T) {
	// Arrange
	dataInput := []TestData{
		{input: "What is 9 squared?", expected: 0, errorMessage: "unsupported operations"},
		{input: "What is 7 cubed?", expected: 0, errorMessage: "unsupported operations"},
		{input: "What is 25 divided by 0?", expected: 0, errorMessage: "division by 0 not allowed"},
	}

	sut := NewEvaluationService()

	for _, v := range dataInput {
		// Act
		result, err := sut.CalculateExpression(v.input)

		// Assert
		if result != v.expected && err == nil {
			t.Errorf("Not correct value of: %v, expexted %v", result, v.expected)
		}
		if err.Error() != v.errorMessage {
			t.Errorf("Not correct value of: %v", err.Error())
		}

	}
}

func TestCalculateExpressionWithNotCorrectInput(t *testing.T) {
	// Arrange
	dataInput := []TestData{
		{input: "What is 5 plus minus 13?", expected: 0, errorMessage: "expressions with invalid syntax"},
		{input: "What is 7 plus cubed?", expected: 0, errorMessage: "unsupported operations"},
		{input: "Who is this", expected: 0, errorMessage: "non-math questions"},
		{input: "What is this", expected: 0, errorMessage: "non-math questions"},
		{input: "What is 7 plus 7 plus", expected: 0, errorMessage: "expressions with invalid syntax"},
	}

	sut := NewEvaluationService()

	for _, v := range dataInput {

		// Act
		result, err := sut.CalculateExpression(v.input)
		// Assert
		if result != v.expected && err == nil {
			t.Errorf("Not correct value of: %v, expexted %v", result, v.expected)
		}
		if err.Error() != v.errorMessage {
			t.Errorf("Not correct value of: %v", err.Error())
		}

	}
}

func TestValidateExpressionWithCorrectInput(t *testing.T) {
	// Arrange
	dataInput := []TestData{
		{input: "What is 5 plus 13?", valid: true},
		{input: "What is 7 minus 5?", valid: true},
		{input: "What is 6 multiplied by 4?", valid: true},
		{input: "What is 25 divided by 5?", valid: true},
		{input: "What is 25?", valid: true},
		{input: "What is 3 plus 2 multiplied by 3?", valid: true},
	}

	sut := NewEvaluationService()

	for _, v := range dataInput {
		// Act
		result, err := sut.ValidateExpression(v.input)
		// Assert
		if result != v.valid || err != nil {
			t.Errorf("Not correct value of: %v, expexted %v", result, v.expected)
		}
	}
}

func TestValidateExpressionWithUnsupportedInput(t *testing.T) {
	// Arrange
	dataInput := []TestData{
		{input: "What is 9 squared?", valid: false, errorMessage: "unsupported operations"},
		{input: "What is 7 cubed?", valid: false, errorMessage: "unsupported operations"},
		{input: "What is 25 divided by 0?", valid: false, errorMessage: "division by 0 not allowed"},
	}

	sut := NewEvaluationService()

	for _, v := range dataInput {
		// Act
		result, err := sut.ValidateExpression(v.input)
		// Assert
		if result != v.valid && err == nil {
			t.Errorf("Not correct value of: %v, expexted %v", result, v.expected)
		}
		if err.Error() != v.errorMessage {
			t.Errorf("Not correct value of: %v", err.Error())
		}

	}
}

func TestValidateExpressionWithNotCorrectInput(t *testing.T) {
	// Arrange
	dataInput := []TestData{
		{input: "What is 5 plus minus 13?", valid: false, errorMessage: "expressions with invalid syntax"},
		{input: "What is 7 plus cubed?", valid: false, errorMessage: "unsupported operations"},
		{input: "Who is this", valid: false, errorMessage: "non-math questions"},
		{input: "What is this", valid: false, errorMessage: "non-math questions"},
		{input: "What is 7 plus 7 plus", valid: false, errorMessage: "expressions with invalid syntax"},
	}

	sut := NewEvaluationService()

	for _, v := range dataInput {
		// Act
		result, err := sut.ValidateExpression(v.input)
		// Assert
		if result != v.valid && err == nil {
			t.Errorf("Not correct value of: %v, expexted %v", result, v.expected)
		}
		if err.Error() != v.errorMessage {
			t.Errorf("Not correct value of: %v", err.Error())
		}

	}
}

func TestSplitInputExpressionValidInut(t *testing.T) {
	// Arrange
	testInput := []TestData{
		{input: "What is 12?", expected: 1},
	}

	for _, v := range testInput {
		// Act
		n, e := SplitInputExpression(v.input)

		// Assert
		if e != nil {

			t.Error("Not correct")
		}

		if len(n) == v.expected {
			t.Log("expected result", n)
		}
	}
}

func TestSplitInputExpressionWithNotCorrectData(t *testing.T) {
	// Arrange
	testInput := []TestData{
		{input: "", expected: 0, errorMessage: "expression inpout can not be empty"},
	}

	for _, v := range testInput {
		// Act
		_, e := SplitInputExpression(v.input)

		// Assert
		if e.Error() != v.errorMessage {
			t.Error("Not correct")
		}

	}
}

func TestExtractNumberWithCorrectData(t *testing.T) {
	// Arrange
	testInput := []TestData{
		{input: "5", expected: 5},
		{input: "-1", expected: -1},
		{input: "0", expected: 0},
	}

	for _, v := range testInput {
		// Act
		n, e := ExtractNumber(v.input)
		// Assert
		if e != nil || n != v.expected {
			fmt.Print(n)
			t.Errorf("Not correct: %v, expected: %v", n, v.expected)
		} else {
			t.Log("the number is correct: ", n)
		}
	}
}

func TestExtractNumberWithNotCorrectData(t *testing.T) {
	// Arrange
	testInput := []TestData{
		{input: "asd", expected: 0},
		{input: "a-a", expected: 0},
	}

	for _, v := range testInput {
		// Act
		n, e := ExtractNumber(v.input)

		// Assert
		if e == nil && n != v.expected {
			fmt.Println(e.Error())
			t.Errorf("Not correct: %v, expected: %v", n, v.expected)
		} else {
			t.Log("the number is correct: ", n)
		}
	}
}

func TestExtractAllNumbersWithCorrectData(t *testing.T) {

	// Arrange
	testInput := []string{
		"32", "123", "-43", "asd",
	}

	testOutput := []int{
		32, 123, -43,
	}

	// Act
	result := ExtractAllNumbers(testInput)

	// Assert
	for i, _ := range testOutput {
		if result[i] != testOutput[i] {
			t.Errorf("Not correct: %v, expected: %v", result[i], testOutput[i])
		}
	}
}

func TestExtractAllNumbersWithNotCorrectData(t *testing.T) {
	// Arrange
	testInput := []string{
		"dasd", "--", "dasada", "asd",
	}

	// Act
	result := ExtractAllNumbers(testInput)

	// Asert
	if len(result) > 0 {
		t.Error("the result must be empty")
	}
}

func TestExtractAllActionsWithCorrectData(t *testing.T) {
	// Arrange
	testInput := []string{
		"plus", "minus", "-43", "asd", "multiplied",
	}

	testOutput := []string{
		"plus", "minus", "multiplied",
	}
	var allowedActions = []string{"plus", "minus", "multiplied", "divided"}

	// Act
	result := ExtractAllActions(testInput, allowedActions)

	// Assert
	for i, _ := range testOutput {
		if result[i] != testOutput[i] {
			t.Errorf("Not correct: %v, expected: %v", result[i], testOutput[i])
		}
	}
}

func TestExtractAllActionsWithNotCorrectData(t *testing.T) {
	// Arrange
	testInput := []string{
		"dasd", "--", "dasada", "asd",
	}
	var allowedActions = []string{"plus", "minus", "multiplied", "divided"}

	// Act
	result := ExtractAllActions(testInput, allowedActions)

	// Assert
	if len(result) > 0 {
		t.Error("the result must be empty")
	}
}

type TestInputData struct {
	inputNumbers []int
	inputActions []string
	expected     int
}

func TestConstructActionsWithCorrectData(t *testing.T) {

	// Arrange
	inputData := []TestInputData{
		{
			inputNumbers: []int{12, 3, 5},
			inputActions: []string{"plus", "minus"},
			expected:     10,
		},
		{
			inputNumbers: []int{2, 3, 5},
			inputActions: []string{"plus", "multiplied"},
			expected:     25,
		},
		{
			inputNumbers: []int{10, 3, 5},
			inputActions: []string{"multiplied", "divided"},
			expected:     6,
		},
		{
			inputNumbers: []int{10, 3},
			inputActions: []string{"multiplied"},
			expected:     30,
		},
	}

	for _, v := range inputData {
		// Act
		result, err := ConstructAndRunActions(v.inputNumbers, v.inputActions)
		// Assert
		if result != v.expected && err == nil {
			t.Errorf("not correct: %v, expected %v", result, v.expected)
		}
	}
}
