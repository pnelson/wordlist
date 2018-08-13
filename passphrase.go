package wordlist

import (
	"math/big"
	"strings"
)

// Passphrase represents the configuration for passphrase generation.
type Passphrase struct {
	count  int
	length int
	sep    string
}

// NewPassphrase returns a new passphrase.
func NewPassphrase(opts ...PassphraseOption) string {
	p := NewPassphraseGenerator(opts...)
	return p.String()
}

// NewPassphraseGenerator returns a new passphrase generator.
func NewPassphraseGenerator(opts ...PassphraseOption) Passphrase {
	p := Passphrase{
		count:  defaultWordCount,
		length: defaultWordLength,
		sep:    defaultWordSeparator,
	}
	for _, option := range opts {
		option(&p)
	}
	return p
}

// String implements the fmt.Stringer interface.
func (p Passphrase) String() string {
	s := make([]string, p.count)
	max := big.NewInt(int64(len(eff)))
	for i := range s {
		for len(s[i]) < p.length {
			s[i] = eff[random(max)]
		}
	}
	return strings.Join(s, p.sep)
}

// PassphraseOption represents a functional option for configuration.
type PassphraseOption func(p *Passphrase)

// defaultWordCount represents the default passphrase
// generation word count.
const defaultWordCount = 6

// WithWordCount sets the passphrase generation word count.
// Defaults to 6.
func WithWordCount(count int) PassphraseOption {
	return func(p *Passphrase) {
		p.count = count
	}
}

// defaultWordLength represents the default passphrase
// generation minimum word length.
const defaultWordLength = 3

// WithWordLength sets the passphrase generation minimum word length.
// Defaults to 3. Limits to 9.
func WithWordLength(length int) PassphraseOption {
	return func(p *Passphrase) {
		if length < maxWordLength {
			p.length = length
		} else {
			p.length = maxWordLength
		}
	}
}

// defaultWordSeparator represents the default passphrase
// generation word separator.
const defaultWordSeparator = " "

// WithWordSeparator sets the passphrase generation word separator.
// Defaults to a single space character.
func WithWordSeparator(sep string) PassphraseOption {
	return func(p *Passphrase) {
		p.sep = sep
	}
}
