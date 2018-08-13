package wordlist

import (
	"testing"
)

func TestNewPassword(t *testing.T) {
	a := NewPassword()
	for i := 0; i < 10; i++ {
		b := NewPassword()
		if b == a {
			t.Fatal("should not generate duplicate password")
		}
	}
}
