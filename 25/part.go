package main

import (
	"fmt"
)

const (
	MODUS    = 20201227
	CARD_KEY = 14222596
	DOOR_KEY = 4057428
)

func generateKey(goal, subjectNumber int) int {
	loops := 0
	value := 1
	for value != goal {
		loops++
		value = ((value % MODUS) * (subjectNumber % MODUS)) % MODUS
	}
	return loops
}

func generateEncryptionKey(key, loops int) int {
	encryptionKey := 1
	for i := 0; i < loops; i++ {
		encryptionKey = ((encryptionKey % MODUS) * (key % MODUS)) % MODUS
	}
	return encryptionKey
}

func main() {
	fmt.Println(generateEncryptionKey(CARD_KEY, generateKey(DOOR_KEY, 7)))
}
