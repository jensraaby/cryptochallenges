package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// TestHexToBase64 uses the test case from the challenge
func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	// Convert hex string input to byte representation (as []bytes)
	hex, err := hex.DecodeString(input)
	if err != nil {
		t.Fail()
	}
	b64 := HexToBase64(hex)
	if output != string(b64) {
		t.Fail()
	}
}

func TestFixedXOR(t *testing.T) {
	leftinput := "1c0111001f010100061a024b53535009181c"
	rightinput := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	left, err := hex.DecodeString(leftinput)
	if err != nil {
		t.Fail()
	}
	right, err := hex.DecodeString(rightinput)
	if err != nil {
		t.Fail()
	}
	result, err := FixedXOR(left, right)
	if err != nil {
		t.Errorf("FixedXOR returned an error")
	}
	hexres := fmt.Sprintf("%x", result)
	if hexres != expected {
		t.Errorf("Wrong result: %s, expected %s", hexres, expected)
	}
}
