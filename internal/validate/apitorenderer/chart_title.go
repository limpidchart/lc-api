package apitorenderer

import "fmt"

const titleMaxLen = 1024

// ErrTitleMaxLen contains error message about max title len.
var ErrTitleMaxLen = fmt.Errorf("title max len is %d", titleMaxLen)

// ValidateChartTitle validates the provided title.
func ValidateChartTitle(title string) error {
	if len(title) > titleMaxLen {
		return ErrTitleMaxLen
	}

	return nil
}
