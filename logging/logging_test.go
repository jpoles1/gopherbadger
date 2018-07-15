package logging

import (
	"errors"
	"testing"
)

func TestErrCheck(t *testing.T) {
	Success("Testing success logging function!")
	Error("Testing error logging function!", errors.New("Test error"))
	Fatal("Testing fatal error logging function!", nil)
}
