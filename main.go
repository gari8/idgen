package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	defaultLength = 16
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_+"
	errText       = "cannot generate id"
)

func main() {
	length := defaultLength
	flag.IntVar(&length, "l", defaultLength, "add log mode")
	flag.Parse()

	id, err := generateID(length)
	if err != nil {
		log.Fatalf("[ERROR] %v\n", err)
	}
	fmt.Println(id)
}

func generateID(length int) (string, error) {
	if length == 0 {
		return "", errors.New(errText)
	}
	buf := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range buf {
		rNum := r.Intn(len(letters))
		buf[i] = letters[rNum]
	}
	return string(buf), nil
}
