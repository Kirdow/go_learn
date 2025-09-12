package main

import (
	"fmt"
	"crypto/sha256"
	"encoding/binary"
	"math/rand/v2"
)

var firstSection bool = true
func Section(title string, fn func()) {
	if !firstSection {
		var i = 2
		for i > 0 {
			i -= 1
			fmt.Println()
		}
	}

	fmt.Printf("] >> %s\n", title)
	firstSection = false

	fn()
}

func MyHash(n uint) uint {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(n))
	hash := sha256.Sum256(buf)
	return uint(binary.BigEndian.Uint64(hash[:8]))
}

func MyRandInt() uint {
	return rand.Uint()
}

func MyDifficultHash(data uint, bits uint) (uint, bool) {
	hash := MyHash(data)
	match := hash & ((1 << (64 - bits)) - 1)
	valid := (match ^ hash) == 0

	return hash, valid
}
