package wordlist

import (
	"math/big"
	"strings"
)

// Password represents the configuration for password generation.
type Password struct {
	length  int
	charset string
}

// NewPassword returns a new password.
func NewPassword(opts ...PasswordOption) string {
	p := NewPasswordGenerator(opts...)
	return p.String()
}

// NewPasswordGenerator returns a new password generator.
func NewPasswordGenerator(opts ...PasswordOption) Password {
	p := Password{
		length:  defaultLength,
		charset: defaultCharset,
	}
	for _, option := range opts {
		option(&p)
	}
	return p
}

// String implements the fmt.Stringer interface.
func (p Password) String() string {
	b := make([]byte, p.length)
	max := big.NewInt(int64(len(p.charset)))
	for i := range b {
		b[i] = p.charset[random(max)]
	}
	return string(b)
}

// PasswordOption represents a functional option for configuration.
type PasswordOption func(p *Password)

// defaultLength represents the default password generation length.
const defaultLength = 12

// WithLength sets the password generation length.
// Defaults to 12.
func WithLength(length int) PasswordOption {
	return func(p *Password) {
		p.length = length
	}
}

// defaultCharset represents the default password
// generation character set of alphanumeric characters.
const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// WithCharset sets the password generation character set.
// Defaults to alphanumeric characters.
func WithCharset(charset string) PasswordOption {
	return func(p *Password) {
		p.charset = charset
	}
}

// WithCharsetFilter sets the exclusion of characters from the charset.
// Apply this option after WithCharset.
func WithCharsetFilter(exclude func(r rune) bool) PasswordOption {
	return func(p *Password) {
		p.charset = strings.Map(func(r rune) rune {
			if exclude(r) {
				return -1
			}
			return r
		}, p.charset)
	}
}
