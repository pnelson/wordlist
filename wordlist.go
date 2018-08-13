// Package wordlist implements secure password and passphrase generators.
package wordlist

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

//go:generate go run wordlist_eff_gen.go

func random(max *big.Int) int {
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		// Cryptographic pseudo-random number generation shouldn't fail, but
		// if it does it is probably worth the panic.
		panic(fmt.Sprintf("wordlist: %v", err))
	}
	return int(n.Int64())
}
