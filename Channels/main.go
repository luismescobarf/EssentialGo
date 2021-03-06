package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var (
	// dataCh   = make(chan Person)
	idx      = 0
	maxCount = 20
)

var dataCh chan Person

func main() {
	// dataCh = make(chan Person)

	// var wg sync.WaitGroup
	// startT := time.Now()
	// fmt.Println("Start deal with data")

	// consumeInts(&wg)
	// produceInts()

	// wg.Wait()
	// fmt.Println("End deal with ", idx, " data")
	// endT := time.Since(startT)
	// fmt.Println("Run time: ", endT)
	// time.Sleep(time.Second)

	// testBoring2()

	// testFakeSearch()
	// testFirstSearch()

	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i * 3
		}
	}()
	go func() {
		for i := 0; ; i++ {
			ch <- i * 5
		}
	}()
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

func produceInts() {

	for {
		if idx >= maxCount { // define when to break the for loop
			break
		}

		person := Person{
			Name: "zhu" + strconv.Itoa(idx),
			Desc: "staff" + strconv.Itoa(idx*idx),
		}
		dataCh <- person
		idx++
	}
	person := Person{
		Name: "zhu" + strconv.Itoa(idx),
		Desc: "staff" + strconv.Itoa(idx*idx),
	}
	dataCh <- person
	close(dataCh)
}

func consumeInts(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range dataCh {
				doPerson(p)
			}
		}()
	}

}

func doPerson(p Person) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Current Person is:", p)
}

// Person :
type Person struct {
	Name string
	Desc string
}
