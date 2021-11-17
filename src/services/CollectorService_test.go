package services

import "testing"

func TestGetErrorsByExpressionWithCorrectData(t *testing.T) {
	// Arrange
	testInput := []TestData{
		{input: "wrong count", expected: 1, errorMessage: "not supported"},
		{input: "wrong count", expected: 2, errorMessage: "not supported"},
		{input: "new wrong count", expected: 1, errorMessage: "not supported"},
		{input: "wrong count", expected: 3, errorMessage: "not supported"},
		{input: "new wrong count", expected: 2, errorMessage: "not supported"},
	}

	sut := NewCollectService()

	for _, v := range testInput {
		sut.AddError(v.input, v.errorMessage, "/validate")
		key := CreateKey(v.input, v.errorMessage, "/validate")

		// Act
		recieved := sut.GetErrorsByExpression(key)

		// Assert
		if recieved.Frequency != v.expected {
			t.Errorf("Not correct value: %v, expected %v", recieved, v.expected)
		} else {
			t.Logf("Correct value: %v", v.expected)
		}
	}
}
func TestGetErrorsByExpressionWithNotExistingKey(t *testing.T) {
	// Arrange
	testInput := []TestData{
		{input: "wrong count", expected: 1, errorMessage: "not supported"},
	}

	sut := NewCollectService()

	for _, v := range testInput {
		sut.AddError(v.input, v.errorMessage, "/validate")
		key := "not-correct-key"

		// Act
		recieved := sut.GetErrorsByExpression(key)

		// Assert
		if recieved.Frequency != 0 || recieved.Expression != "" {
			t.Errorf("Not correct value: %v, expected %v", recieved, v.expected)
		} else {
			t.Logf("Correct value: %v", v.expected)
		}
	}
}

func TestAddError(t *testing.T) {
	// Arrange
	testInput := []TestData{
		{input: "wrong count", expected: 1, errorMessage: "not supported"},
		{input: "wrong count", expected: 2, errorMessage: "not supported"},
		{input: "new wrong count", expected: 1, errorMessage: "not supported"},
		{input: "wrong count", expected: 3, errorMessage: "not supported"},
		{input: "new wrong count", expected: 2, errorMessage: "not supported"},
	}

	testOutput := []int{3, 2}

	sut := NewCollectService()

	for _, v := range testInput {
		sut.AddError(v.input, v.errorMessage, "/validate")
	}

	// Act
	recieved := sut.GetAllErrors()

	// Assert
	if len(recieved) != 2 {
		t.Error("Not correct count of errors: ", len(recieved))
	}
	for i := range recieved {
		frequencyExpected := testOutput[i]
		freqRecieved := recieved[i].Frequency
		if frequencyExpected != freqRecieved {
			t.Errorf("Not correct value: %v, expected %v", recieved[i].Frequency, frequencyExpected)
		} else {
			t.Logf("Correct value: %v", freqRecieved)
		}
	}

}
