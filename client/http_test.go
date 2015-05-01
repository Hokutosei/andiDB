package main

import (
	"fmt"
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randSeq generate random KEY string
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// TestPost will post request to server
func TestPost(t *testing.T) {
	for i := 0; i < 200; i++ {
		fmt.Sprintf("h: %d", i)

		values := []string{randSeq(20)}
		Post("set", randSeq(10), values)
	}
}
