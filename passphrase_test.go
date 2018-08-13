package wordlist

import "testing"

func TestNewPassphrase(t *testing.T) {
	a := NewPassphrase()
	for i := 0; i < 10; i++ {
		b := NewPassphrase()
		if b == a {
			t.Fatal("should not generate duplicate passphrase")
		}
	}
}
