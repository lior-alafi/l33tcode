package models

import (
	"fmt"
	"strings"
)

func IsEmpty(s string, field string) error {
	if strings.Trim(s, " \t\n") == "" {
		return fmt.Errorf("field: %s must be filled", field)
	}

	return nil
}
