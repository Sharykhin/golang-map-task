package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type NumBucket map[int64]int64

type Value struct {
	IsComplex bool
	Val int64
}

func calculate(values []Value) (NumBucket, error) {
	b := make(NumBucket)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, v := range values {
		if v.IsComplex {
			wg.Add(1)
			go func(val int64){
				defer wg.Done()
				time.Sleep(1*time.Second)
				mu.Lock()
				b[val]++
				mu.Unlock()
			}(v.Val)
		} else {
			mu.Lock()
			b[v.Val]++
			mu.Unlock()
		}
	}

	wg.Wait()

	return b, nil
}

func main() {

	res, err := calculate([]Value{
		{false, 2},
		{false, 2},
		{true, 2},
		{false, 1},
		{false, 3},
		{true, 4},
		{true, 4},
		{true, 5},
		{true, 5},
		{true, 5},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
