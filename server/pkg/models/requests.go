package models

type SelectCodeExecuterRequest struct {
	Name string
}

type CodeSubmitRequest struct {
	QID               string
	Language          string
	Code              string
	StartTimestamp    string
	User              string
	RecievedTimestamp string
}

//TODO: support pagination
// type RequestPagination struct {
// 	PageNum  int
// 	PageSize int
// 	Marker   string
// }
// type ListLanguagePagination struct {
// 	PageNum  int
// 	PageSize int

// 	Marker string

// 	Languages []Language
// }

// type ListQuestionPagination struct {
// 	PageNum  int
// 	PageSize int

// 	Marker string

// 	Languages []Language
// }
