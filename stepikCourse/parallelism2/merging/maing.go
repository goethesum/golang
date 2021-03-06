package main

import (
	// пакет используется для проверки выполнения условия задачи, не удаляйте его

	"sync"
)

var lock sync.Mutex

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
		lock.Lock()

		f1res := make([]*int, n)
		f2res := make([]*int, n)

		input := func(input <-chan int, results []*int) {
			for i := 0; i < n; i++ {
				x := <-input
				go func(i int, x int) {
					res := fn(x)
					results[i] = &res
				}(i, x)
			}
		}

		go input(in1, f1res)
		go input(in2, f2res)

		go func() {
			i := 0
			for true {
				if f1res[i] != nil && f2res[i] != nil {
					res := *f1res[i] + *f2res[i]
					out <- res
					if i++; i == n {
						lock.Unlock()
						return
					}
				}
			}
		}()

	}(fn, in1, in2, out, n)
}

func main() {

}
