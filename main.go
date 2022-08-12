package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"time"
)

type Stats struct {
	hitpoints int
	defense   int
	energy    int
}

type Character struct {
	stats Stats
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func main() {
	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := "This is a secret"

	c := EncryptAES([]byte(key), pt)

	// plaintext
	fmt.Println(pt)

	// ciphertext
	fmt.Println(c)

	// decrypt
	DecryptAES([]byte(key), c)
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

func timeTesting() {
	now := time.Now()

	loc, _ := time.LoadLocation("Europe/Oslo")
	fmt.Printf("Oslo Time: %s\n", now.In(loc))
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func closureTest() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	test := nextInt()

	fmt.Println("test: ", test)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
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

func sliceTest() {
	slice := make([]string, 4)

	fmt.Println("Slice: ", slice)

	fmt.Println("len(slice): ", len(slice))

	slice[0] = "Hello"
	slice[1] = ", "
	slice[2] = "world"
	slice[3] = "!"

	fmt.Println("Slice: ", slice)

	slice = append(slice, "appended")

	fmt.Println("Slice: ", slice)

	sliceCopy := make([]string, len(slice))

	copy(sliceCopy, slice)

	fmt.Println("SliceCopy: ", sliceCopy)

	fmt.Println("Slice[:]: ", slice[:])

	capTest := slice[1:3]

	fmt.Println("capTest: ", capTest)
	fmt.Println("cap(capTest): ", cap(capTest))

	capTest = capTest[:cap(capTest)]

	fmt.Println("capTest: ", capTest)
	fmt.Println("cap(capTest): ", cap(capTest))

	s := []string{"growing"}

	fmt.Println("S: ", s)
	fmt.Println("cap(s): ", cap(s))

	t := make([]string, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0

	fmt.Println("T: ", t)
	fmt.Println("cap(t): ", cap(t))

	copy(t, s)

	fmt.Println("T: ", t)

	s = t

	fmt.Println("S: ", s)
	fmt.Println("cap(s): ", cap(s))
}

func arrayTest() {
	var a [5]int

	fmt.Println("a: ", a)

	a[2] = 7357

	fmt.Println("a[2]: ", a[2])

	b := [2]int{2, 3}

	fmt.Println("len(b): ", len(b))
	fmt.Println("b: ", b)

	c := [...]string{"This", "is", "a", "test"}

	fmt.Println("c: ", c)
}

func mapTest() {
	test := make(map[string]int)

	fmt.Println("Test: ", test)

	test["One"] = 1

	fmt.Println("Test: ", test)

	test["Two"] = 2

	fmt.Println("Test: ", test)

	fmt.Println("Test[One] ", test["One"])

	fmt.Println("len(test): ", len(test))

	test["Delete"] = 631373

	fmt.Println("Test: ", test)

	delete(test, "Delete")

	fmt.Println("Test: ", test)

	value, exists := test["Two"]

	fmt.Println("Value: ", value)
	fmt.Println("Exists: ", exists)
}

func switchTest() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool: %t\n", i)
		case int:
			fmt.Printf("I'm an int: %d\n", i)
		default:
			fmt.Printf("Don't know type %T: %s\n", t, i)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
