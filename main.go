package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

const slice_size = 20

func readFile(f string) *string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text string

	for scanner.Scan() {
		text = scanner.Text()
	}

	err = file.Close()
	if err != nil {
		log.Println(err)
	}

	return &text
}

func countChar(c byte, str string, ch chan int, id int, snc *sync.WaitGroup) {
	cnt := 0
	for i := 0; i < len(str); i++ {
		if (str)[i] == c {
			cnt++
		}
	}
	ch <- cnt
	defer snc.Done()
	fmt.Println("id:", id, " total number of ", c, "(s) found in this go routine  -> ", cnt)
}

func main() {
	fileName := "test2.txt"

	counterChannel := make(chan int)
	defer close(counterChannel)
	var count int

	input := readFile(fileName)
	id := 1
	group := sync.WaitGroup{}
	for i := 0; i < len(*input); i += slice_size {
		start := i
		end := i + slice_size
		if end > len(*input) {
			end = len(*input)
		}
		group.Add(1)
		go countChar('a', (*input)[start:end], counterChannel, id, &group)
		count += <-counterChannel
		id++
	}

	group.Wait()

	fmt.Println("Total Number of character found : ", count)

}
