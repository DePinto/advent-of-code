package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var users = []string{
	"Kramer",
	"Eric",
	"Scotty",
	"Danny",
	"Jeff",
	"Mike",
	"Albert",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Printf("%s\n", strings.Join(shuffle(users), "\n"))
}

func shuffle(vs []string) []string {
	xs := make([]string, len(vs))
	copy(xs, vs)

	for i := range xs {
		j := rand.Intn(len(xs))
		xs[i], xs[j] = xs[j], xs[i]
	}

	return xs
}
