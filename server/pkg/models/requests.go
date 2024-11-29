package models

type CodeSubmitRequest struct {
	QID               string
	Language          string
	Code              string
	StartTimestamp    string
	User              string
	RecievedTimestamp string
}
