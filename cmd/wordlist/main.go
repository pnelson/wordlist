// wordlist generates passwords and passphrases.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/pnelson/wordlist"
)

var (
	h = flag.Bool("h", false, "show this usage information")

	n = flag.Int("n", 1, "number of password/passphrase generations")

	wordSep    = flag.String("sep", " ", "passphrase word separator")
	wordCount  = flag.Int("word-count", 6, "passphrase word count")
	wordLength = flag.Int("word-length", 3, "minimum passphrase word length (limit 9)")

	pw      = flag.Bool("pw", false, "generate password instead of passphrase")
	length  = flag.Int("length", 12, "password length")
	charset = flag.String("charset", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "password charset")
	noLower = flag.Bool("no-lower", false, "passwords without lowercase characters")
	noUpper = flag.Bool("no-upper", false, "passwords without uppercase characters")
	noDigit = flag.Bool("no-digit", false, "passwords without digits")
)

func init() {
	log.SetFlags(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]...\n\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if *h {
		flag.Usage()
		return
	}
	if *n < 1 {
		*n = 1
	}
	if *pw {
		options := []wordlist.PasswordOption{
			wordlist.WithLength(*length),
			wordlist.WithCharset(*charset),
		}
		if *noLower {
			options = append(options, wordlist.WithCharsetFilter(unicode.IsLower))
		}
		if *noUpper {
			options = append(options, wordlist.WithCharsetFilter(unicode.IsUpper))
		}
		if *noDigit {
			options = append(options, wordlist.WithCharsetFilter(unicode.IsDigit))
		}
		for i := 0; i < *n; i++ {
			password := wordlist.NewPassword(options...)
			fmt.Println(password)
		}
		return
	}
	options := []wordlist.PassphraseOption{
		wordlist.WithWordCount(*wordCount),
		wordlist.WithWordLength(*wordLength),
		wordlist.WithWordSeparator(*wordSep),
	}
	for i := 0; i < *n; i++ {
		passphrase := wordlist.NewPassphrase(options...)
		fmt.Println(passphrase)
	}
}
