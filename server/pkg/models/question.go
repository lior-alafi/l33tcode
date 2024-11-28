package models

type Test struct {
	Inputs   string //e.g "nums: [1,2,3,4] letter: 'd'"
	Expected string
}
type SupportedLanguage struct {
	Language  string
	Prototype string
}
type Question struct {
	Id                  string
	Title               string
	Description         string
	SupportedLanguagges []string
	Tests               []Test
}
