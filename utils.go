package main

import (
	"math/rand"
	"time"
)

const numChars = 6
const numVariants = 62

func genShort() string {
	runes := make([]rune, numChars)
	rand.Seed(time.Now().Unix())
	data := getPreparedData()
	for i := 0; i < numChars; i++ {
		runes[i] = data[rand.Intn(numVariants)]
	}

	return string(runes)
}

func getPreparedData() []rune {
	res := make([]rune, numVariants)
	top := 0

	rng, count := genRange('a', 'z')
	top += count
	copy(res[:count], rng)

	rng, count = genRange('A', 'Z')
	copy(res[top:top+count], rng)
	top += count

	rng, count = genRange('0', '9')
	copy(res[top:top+count], rng)

	return res
}

func genRange(start, finish rune) ([]rune, int) {
	count := int(finish - start) + 1
	res := make([]rune, count)
	for c := 0; c < count; c++ {
		res[c] = start + rune(c)
	}
	return res, count
}
