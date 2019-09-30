package funcerr

import (
	"some-rest-api/logger"
)

func isError(err error) bool {
	if err == nil {
		return true
	}
	return false
}

/*
IsErrored := function to see if there are any errors upon performing an operation.
returns true if error is present and false if err is nil
*/
func IsErrored(err error, msg string) bool {
	if !isError(err) {
		logger.Message(msg, "error")
		return true
	}
	return false
}

/*
IsErroredWithLevel := works similar to that of IsErrored, the only
difference being that the function accepts a level for the logs as well.
*/
func IsErroredWithLevel(err error, msg string, msgType string) bool {
	if !isError(err) {
		logger.Message(msg, msgType)
		return true
	}
	return false
}
