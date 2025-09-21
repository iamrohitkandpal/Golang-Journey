package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"
)

// Sending
// func processNum (numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing Number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	numChan := make(chan int)

// 	go processNum(numChan)

// 	for {
// 		numChan <- rand.Intn(100)
// 	}
// }

// Recieving
// func sum(res chan int, n1 int, n2 int) {
// 	numRes := n1 + n2
// 	res <- numRes
// }

// func main() {
// 	res := make(chan int)

// 	go sum(res, 2, 3)

// 	result := <- res
// 	fmt.Println(result)
// }

// Synchronization
// func task(done chan bool) {
// 	defer func ()  {
// 		done <- true
// 	} ()

// 	fmt.Println("Processing....")
// }

// func main() {
// 	done := make(chan bool)

// 	go task(done)

// 	<- done
// }

// Buffer Channel
func emailSender(emailChan chan string, done chan bool) {
	defer func ()  {
		done <- true
	}()

	for email := range emailChan {
		fmt.Println("Sending Mail to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 100)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("Done Adding Emails....")

	go emailSender(emailChan, done)

	close(emailChan)			// Highly Important thing
	<- done
}