package response

type ValidatedMathExpression struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason"`
}
