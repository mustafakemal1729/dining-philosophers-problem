package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	// number of philosophers
	N = 5

	// number of eats
	E = 100

	// Sleep range for sleepRand function
	sleepRandMin = 1
	sleepRandMax = 10
)

var wg sync.WaitGroup

func sleepRand() {
	r := rand.Int31n(sleepRandMax-sleepRandMin+1) + sleepRandMin
	time.Sleep(time.Millisecond * time.Duration(r))
}

func takeLeftFork(number int, forkMutex *sync.Mutex) {
	sleepRand()
	forkMutex.Lock()
	fmt.Printf("Philosopher %v has taken the left fork %v\n", number, number)
}

func takeRightFork(number int, forkMutex *sync.Mutex) {
	sleepRand()
	forkMutex.Lock()
	fmt.Printf("Philosopher %v has taken the right fork %v\n", number, (number+1)%N)
}

func eatAndRelease(number int, leftForkMutex, rightForkMutex *sync.Mutex) {
	sleepRand()
	fmt.Printf("Philosopher %d has eaten\n", number)
	rightForkMutex.Unlock()
	leftForkMutex.Unlock()
	fmt.Printf("Philosopher %d has released the forks\n", number)
}

func philosopher(number int, forks *[N]sync.Mutex) {
	defer wg.Done()
	for i := 0; i < E; i++ {
		left := number
		right := (left + 1) % N

		// Ensure that philosophers always pick up the lower-numbered fork first
		if left < right {
			takeLeftFork(number, &forks[left])
			takeRightFork(number, &forks[right])
		} else {
			takeRightFork(number, &forks[right])
			takeLeftFork(number, &forks[left])
		}

		eatAndRelease(number, &forks[left], &forks[right])
	}
}

func main() {
	forks := [N]sync.Mutex{}
	fmt.Printf("%v philosophers have begun to eat\n", N)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go philosopher(i, &forks)
	}
	wg.Wait()
}
