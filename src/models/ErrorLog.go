package models

type ErrorLog struct {
	Expression string
	Endpoint   string
	Frequency  int
	Types      string
}
