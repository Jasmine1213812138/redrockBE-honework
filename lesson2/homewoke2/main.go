package main

import (
	"fmt"
	"sync"
	"time"
)

func download(filename string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(time.Second)
	results <- fmt.Sprintf("%s 下载完成", filename)
}
func main() {
	file := []string{"file1.zip", "file2.zip", "file3.zip"}
	ch := make(chan string, len(file))
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download(file[i], &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for result := range ch {
		fmt.Println(result)
	}

}
