package main

import (
	"crypto/aes"
	"encoding/hex"
	"example/learner-api/api"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(api.IndexHandler),
	}

	server.ListenAndServe()
}

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	CheckError(err)

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CopyDigits(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")

	b, _ := ioutil.ReadFile(filename)

	fmt.Println("b: ", b)

	b = digitRegexp.Find(b)

	fmt.Println("b: ", b)

	c := make([]byte, len(b))
	copy(c, b)

	fmt.Println("c: ", c)

	c[0] = 0
	d := c
	d[0] = 99

	fmt.Println("c: ", c)
	fmt.Println("b: ", b)
	fmt.Println("d: ", d)

	return b
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice

		fmt.Println("slice: ", slice)
	}
	slice = slice[0:n]

	fmt.Println("slice: ", slice)

	copy(slice[m:n], data)
	return slice
}
