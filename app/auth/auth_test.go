package auth

import (
	"errors"
	"testing"
)

func TestAuth(t *testing.T) {
	a := 1

	if a != 1 {
		t.Error(errors.New("Testing Auth"))
	}
}
