package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {

	testdata := make(map[string][]byte)
	strings := []string{"hello world", "Hello World", "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"}

	for _, v := range strings {
		testdata[v] = []byte(fmt.Sprintf("%x", v))
	}

	// Test the XOR code:
	diff, err := FixedXOR(testdata[strings[0]], testdata[strings[1]])
	if err != nil {
		fmt.Fprintf(os.Stderr, "XOR failed")
	}
	fmt.Println(diff)

	Ex3()
	Ex4()
}
func HexToBase64(in []byte) []byte {

	// We hold the encoded stream in a byte buffer
	var w bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &w)
	encoder.Write(in)
	encoder.Close()
	return w.Bytes()
}

func FixedXOR(left, right []byte) ([]byte, error) {
	// take each element and XOR to produce the result.
	out := make([]byte, len(left))
	// first check equal length slices
	if len(left) != len(right) {
		return out, fmt.Errorf("Cannot XOR different length streams (%v and %v)", len(left), len(right))
	}

	// I am sure there is some fast way of doing this in C!
	for i, v := range left {
		out[i] = v ^ right[i]
	}

	return out, nil
}

func Ex3() {
	fmt.Println("Exercise 3: Single-byte XOR cipher")
	// the input has been XOR'ed against a single character
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	asbytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hex string", input)
	fmt.Println("Byte representation:", asbytes, "Length:", len(asbytes))
	fmt.Println("String conversion:", string(asbytes))
	fmt.Println("Byte values of alphabet:")
	// for i := 0; i < 26; i++ {
	// 	val := byte('a') + byte(i)
	// 	fmt.Printf("%v: %v\n", string(val), val)
	// }

	counter := make(map[byte]int)
	for _, v := range asbytes {
		counter[v] += 1
	}
	fmt.Println("Bytes ordered descending by count:")
	descendingBytes := SortedKeys(counter)
	fmt.Println(descendingBytes)
	// create a strip of this byte for my XOR:
	strip := bytes.Repeat([]byte{descendingBytes[0]}, len(asbytes))
	deciphered, err := FixedXOR(strip, asbytes)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println("Deciphered bytes:", deciphered)
	fmt.Println("Deciphered string:", string(deciphered))
}

func Ex4() {

	fmt.Println("============\nEx4: detecting which string is encrypted")
	// this time, we load in 4.txt and put the strings in a [][]string
	var encodedHex []string

	textfile, err := os.Open("4.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer textfile.Close()

	scanner := bufio.NewScanner(textfile)
	for scanner.Scan() {

		encodedHex = append(encodedHex, scanner.Text())
	}
	fmt.Println(len(encodedHex))

}
