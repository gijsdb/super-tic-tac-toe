package models

type ResponseError struct {
	Err        string `json:"Error"`
	Stacktrace string `json:"stacktrace"`
}
