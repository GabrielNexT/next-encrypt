package internal

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
	"lukechampine.com/blake3"
)

func blake256(in []byte) [32]byte {
	return blake3.Sum256(in)
}

func getSalt(password []byte) []byte {
	a := blake256(password)
	b := blake256(a[:])
	c := blake256(b[:])
	return c[:]
}

func getNonce(password []byte) []byte {
	a := blake256(password)
	return a[0:12]
}

func GetKeyAndNonceFromPassword() ([]byte, []byte) {
	fmt.Print("Enter password: ")
	passwd, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	salt := getSalt(passwd)

	return argon2.IDKey(passwd, salt, 1, 64*1024, 4, 32), getNonce(passwd)
}
