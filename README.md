# wordlist

Package wordlist implements secure password and passphrase generators.

## Usage

In the most simple use case, generate a password or passphrase with defaults.

```go
password := wordlist.NewPassword()
// 9pbeFr4VamXq

passphrase := wordlist.NewPassphrase()
// starless curse respect shelter murmuring frying
```

You can get creative with this.

Maybe you need a six digit PIN:

```go
pin := wordlist.NewPassword(wordlist.WithLength(6), wordlist.WithCharset("0123456789"))
// 742351
```

Or an unambiguous eight character discount code:

```go
code := wordlist.NewPassword(wordlist.WithLength(8), wordlist.WithCharset("ABCDEFGHKLMNPRSTWXY3456789"))
// E4PRBPPL
```

Or many unambiguous eight character discount codes:

```go
code := wordlist.NewPasswordGenerator(wordlist.WithLength(8), wordlist.WithCharset("ABCDEFGHKLMNPRSTWXY3456789"))
fmt.Println(code)
fmt.Println(code)
fmt.Println(code)
// 4Y9ELLBT
// LAG7EFSN
// 83PAFPTT
```

## Licenses

Package wordlist is licensed under the terms described in
[LICENSE](https://github.com/pnelson/wordlist/blob/master/LICENSE).

Passphrase generation depends on data generated from
[eff_large_wordlist.txt](https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt)
found in the blog post
[Deep Dive: EFF's New Wordlists for Random Passphrases](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).
