package main

import (
	"fmt"
	///"io/ioutil"
	//"net/http"
	"time"
)

/*func main() {
	intChan := make(chan int)
	for i := 1; i <= 4; i++ {
		go func(j int) {
			time.Sleep(time.Duration(j) * time.Second)
			intChan <- j
			if j == 6 {
				close(intChan)
			}
		}(i)
	}

	for j := range intChan {
		fmt.Printf("%d,", j)
	}
	fmt.Println("Done")
	n := time.Now().UnixNano()
	rand.Seed(7)
	i := rand.Intn(5)
	fmt.Println("n=", n, "  i=", i)
}

func wait(c chan int) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(5)
	time.Sleep(time.Duration(i) * time.Second)
	c <- i
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go wait(chan1)
	go wait(chan2)

	select {
	case i := <-chan1:
		fmt.Printf("Received %d on chan 1", i)
	case i := <-chan2:
		fmt.Printf("Received %d on chan 2", i)
	}
}
func main() {
	responseChan := make(chan *http.Response)
	go getURL(responseChan)

	timer := time.NewTimer(700 * time.Millisecond)

	select {
	case response := <-responseChan:
		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		fmt.Printf("Received response: %s", string(body))

	case <-timer.C:
		fmt.Printf("Request timed out")
	}
}

func getURL(c chan *http.Response) {
	response, err := http.Get("https://www.google.co.nz/search?q=golang")
	if err != nil {
		panic(err)
	}
	c <- response
}*/

func main11() {
	stopChan := make(chan bool)
	go func() {
		time.Sleep(4100 * time.Millisecond)
		stopChan <- true
	}()

	timer := time.NewTimer(time.Second)

LOOP:
	for {
		select {
		case <-timer.C:
			// In reality you’d have some polling code here
			fmt.Println("Tick")
			timer.Reset(time.Second)
		case <-stopChan:
			fmt.Println("Boom")
			break LOOP
		}
	}
}
