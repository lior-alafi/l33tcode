package models

import (
	"fmt"
)

type Test struct {
	Inputs   string //e.g "nums: [1,2,3,4] letter: 'd'"
	Expected string
}
type SupportedLanguage struct {
	Language  string //e.g python
	Call      string // foo(a,b,c,d) or Solution().foo(a,b,c,d)
	Solution  string
	Prototype string
	Tests     []Test
}
type Question struct {
	Id                  string
	Title               string
	Description         string
	SupportedLanguagges []SupportedLanguage

	Owner string
}

func (q *Question) Validate() error {
	if err := IsEmpty(q.Title, "Title"); err != nil {
		return err
	}

	if err := IsEmpty(q.Description, "Description"); err != nil {
		return err
	}

	if err := q.validateSupportedLanguages(); err != nil {
		return err
	}

	return nil
}

func (q *Question) validateSupportedLanguages() error {
	for i, sl := range q.SupportedLanguagges {
		if err := IsEmpty(sl.Language, fmt.Sprintf("SupportedLanguagges[%d].Language", i)); err != nil {
			return err
		}

		if err := IsEmpty(sl.Call, fmt.Sprintf("SupportedLanguagges[%d].Call", i)); err != nil {
			return err
		}

		if err := IsEmpty(sl.Prototype, fmt.Sprintf("SupportedLanguagges[%d].Prototype", i)); err != nil {
			return err
		}
		// TODO: future support for test input
		if err := IsEmpty(sl.Solution, fmt.Sprintf("SupportedLanguagges[%d].Solution", i)); err != nil {
			return err
		}

		if len(sl.Tests) == 0 {
			return fmt.Errorf("supported language: %s must have at least 1 test", sl.Language)
		}
		for j, tst := range sl.Tests {
			if err := IsEmpty(tst.Expected, fmt.Sprintf("SupportedLanguagges[%d].Tests[%d].Expected", i, j)); err != nil {
				return err
			}

			if err := IsEmpty(tst.Inputs, fmt.Sprintf("SupportedLanguagges[%d].Tests[%d].Inputs", i, j)); err != nil {
				return err
			}
		}
	}
	return nil
}
