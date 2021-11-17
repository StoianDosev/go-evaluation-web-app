package services

import (
	"evaluationapp/src/models"
	"fmt"
)

type ICollectorService interface {
	GetErrorsByExpression(expression string) models.ErrorLog
	GetAllErrors() []models.ErrorLog
	AddError(expression string, errorMsg string, endpoint string)
}

type CollectorSerivce struct {
	_errorExpressions map[string]*models.ErrorLog
}

func NewCollectService() *CollectorSerivce {
	return &CollectorSerivce{
		_errorExpressions: make(map[string]*models.ErrorLog),
	}
}

func (c *CollectorSerivce) GetErrorsByExpression(expression string) models.ErrorLog {

	element := c._errorExpressions[expression]

	if element == nil {
		return models.ErrorLog{}
	}
	return *element
}

func (c *CollectorSerivce) GetAllErrors() []models.ErrorLog {
	list := make([]models.ErrorLog, 0, len(c._errorExpressions))

	for _, value := range c._errorExpressions {
		list = append(list, *value)
	}
	return list
}

func (c *CollectorSerivce) AddError(expression string, errorMsg string, endpoint string) {
	var v *models.ErrorLog
	var found bool
	key := CreateKey(expression, errorMsg, endpoint)
	v, found = c._errorExpressions[key]
	if found {
		v.Frequency = v.Frequency + 1
		return
	}

	c._errorExpressions[key] = &models.ErrorLog{
		Expression: expression,
		Endpoint:   endpoint,
		Types:      errorMsg,
		Frequency:  1,
	}
}

func CreateKey(expression string, errorMsg string, endpoint string) string {
	return fmt.Sprintf("%v-%v-%v", endpoint, errorMsg, expression)
}
