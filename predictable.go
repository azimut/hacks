package main

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func seedRandom() {
	rand.Seed(time.Now().Unix())
}
