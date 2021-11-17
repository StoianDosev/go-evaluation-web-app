package response

type ErrorExpression struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	Types      string `json:"types"`
}
