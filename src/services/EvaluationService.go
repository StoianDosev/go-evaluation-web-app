package services

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type IEvaluationService interface {
	ValidateExpression(exp string) (bool, error)
	CalculateExpression(exp string) (int, error)
}

var _allActions = []string{"plus", "minus", "multiplied", "divided", "cubed", "squared"}
var _notAllowedActions = []string{"cubed", "squared"}
var _actionFuncs = map[string]func(int, int) (int, error){
	"plus": add, "minus": substract, "multiplied": multiply, "divided": devide,
}

type EvaluationService struct {
}

func NewEvaluationService() *EvaluationService {
	return &EvaluationService{}
}

func (s *EvaluationService) ValidateExpression(exp string) (bool, error) {
	_, err := s.CalculateExpression(exp)
	if err != nil {

		return false, err
	}
	return true, nil
}

func (s *EvaluationService) CalculateExpression(exp string) (int, error) {
	splitExpression, err := SplitInputExpression(exp)
	if err != nil {
		return 0, err
	}

	numbers := ExtractAllNumbers(splitExpression)
	actions := ExtractAllActions(splitExpression, _allActions)
	result, err := ConstructAndRunActions(numbers, actions)

	return result, err
}

func SplitInputExpression(wholeExpression string) ([]string, error) {

	if wholeExpression != "" {
		return strings.Fields(wholeExpression), nil
	}

	return nil, errors.New("expression inpout can not be empty")
}

func ExtractNumber(expectedNumber string) (int, error) {

	re := regexp.MustCompile("^-?[0-9]+")

	arrNum := re.FindAllString(expectedNumber, -1)
	if len(arrNum) == 1 {
		n, err := strconv.Atoi(arrNum[0])
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	if len(arrNum) == 0 {
		return 0, errors.New("not a number")
	}
	return 0, errors.New("not allowed more than one number")
}

func EvaluatesAction(phraze string, actions map[string]func(int, int) (int, error)) (func(int, int) (int, error), bool) {

	fun, found := actions[phraze]

	return fun, found
}

func add(a int, b int) (int, error) {
	return a + b, nil
}

func substract(a int, b int) (int, error) {
	return a - b, nil
}

func multiply(a int, b int) (int, error) {
	return a * b, nil
}

func devide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by 0 not allowed")
	}
	return a / b, nil
}

func ExtractAllActions(input []string, allowedActions []string) []string {
	var output []string
	for _, v := range input {
		removedQuestion := strings.ReplaceAll(v, "?", "")
		for _, a := range allowedActions {

			if a == removedQuestion {
				output = append(output, v)
			}
		}
	}
	return output
}

func ExtractAllNumbers(input []string) []int {
	var numbers []int
	for _, v := range input {
		n, e := ExtractNumber(v)
		if e == nil {
			numbers = append(numbers, n)
		}
	}

	return numbers
}

func IsActionsAllowed(actions []string, notAllowedActions []string) bool {

	for _, v := range actions {
		removedQuestion := strings.ReplaceAll(v, "?", "")
		for _, a := range notAllowedActions {
			if a == removedQuestion {
				return false
			}
		}
	}
	return true
}

func ValidateInputs(numbers []int, actions []string) error {
	if len(numbers) == 0 && len(actions) == 0 {
		err := errors.New("non-math questions")
		return err
	}

	if len(numbers) <= len(actions) || len(numbers) != 1 && len(actions) == 0 {
		err := errors.New("expressions with invalid syntax")
		return err
	}

	return nil
}

func ConstructAndRunActions(numbers []int, actions []string) (int, error) {
	isAllowed := IsActionsAllowed(actions, _notAllowedActions)
	if !isAllowed {
		err := errors.New("unsupported operations")
		return 0, err
	}

	err := ValidateInputs(numbers, actions)
	if err != nil {
		return 0, err
	}

	if len(numbers) == 1 {
		return numbers[0], nil
	}

	var currResult int
	copyNumbers := make([]int, len(numbers))
	copy(copyNumbers, numbers)
	isFistRun := true
	for _, a := range actions {
		f, _ := EvaluatesAction(a, _actionFuncs)

		for i, _ := range numbers {

			currNumber := numbers[i]
			copyNumbers = copyNumbers[1:] //popfirst

			if i+1 < len(numbers) && isFistRun {
				isFistRun = false
				currN := numbers[i+1]
				currResult = currNumber
				currNumber = currN
				copyNumbers = copyNumbers[1:] // popfirst
			}

			var err error
			currResult, err = f(currResult, currNumber)

			if err != nil {
				return 0, err
			}

			numbers = copyNumbers
			break
		}
	}
	return currResult, nil
}
