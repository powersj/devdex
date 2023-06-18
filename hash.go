package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

// List of crypto algorithms https://pkg.go.dev/crypto

func MD5Sum(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

func SHA1Sum(input string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(input)))
}

func SHA256Sum(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}

func SHA512Sum(input string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(input)))
}

func SHA3_256Sum(input string) string {
	return fmt.Sprintf("%x", sha3.Sum256([]byte(input)))
}

func SHA3_512Sum(input string) string {
	return fmt.Sprintf("%x", sha3.Sum512([]byte(input)))
}

func Blake2b256Sum(input string) string {
	return fmt.Sprintf("%x", blake2b.Sum256([]byte(input)))
}

func Blake2b512Sum(input string) string {
	return fmt.Sprintf("%x", blake2b.Sum512([]byte(input)))
}
