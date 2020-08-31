package common

import "strconv"

type ErrorWithStatus struct {
	Code int
}

func (e *ErrorWithStatus) Error() string {
	return "failed with status " + strconv.Itoa(e.Code)
}
