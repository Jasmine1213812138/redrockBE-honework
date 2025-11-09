package main

import (
	"fmt"
	"sync"
)

type Conter struct {
	mu    sync.Mutex
	count int
}

func (c *Conter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
func (c *Conter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func main() {
	con := Conter{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				con.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println(con.count)
}
