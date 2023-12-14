package main

import (
	"math/rand"
	"sync"
	"time"
)

const (

	// number of philosophers
	N = 5

	// number of eats
	E = 100
)

var wg sync.WaitGroup

func sleep_rand() {
	r := rand.Int31n(10) + 1
	time.Sleep(time.Millisecond * time.Duration(r))
}
