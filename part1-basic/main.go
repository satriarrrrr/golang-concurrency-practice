package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var ch chan bool

// func main() {
// 	fmt.Println("Hello World")

// 	ch = make(chan bool)

// 	go StartAnotherRoutine()
// 	fmt.Println("Waiting for 2 secs")
// 	time.Sleep(2 * time.Second)
// 	<-ch
// 	fmt.Println("Message received")
// 	fmt.Println("Wait for 2 more secs")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Finished !")
// }

// func StartAnotherRoutine() {
// 	fmt.Println("[Another] is going to send ...")
// 	ch <- true

// 	fmt.Println("[Another] has finished sending")
// }

func main() {
	go boring("boring!")
	time.Sleep(2 * time.Second)
	fmt.Println("I'm listening.")
	time.Sleep(30 * time.Millisecond)
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
