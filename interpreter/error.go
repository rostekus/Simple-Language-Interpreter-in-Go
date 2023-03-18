package interpreter

import "fmt"

type NameError struct {
	message string
	line    int
	file    string
}

func (e NameError) Error() string {
	return fmt.Sprintf("NameError: %s\n%s, line %d", e.message, e.file, e.line)
}
