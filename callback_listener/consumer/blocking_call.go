package main

import (
	"fmt"
)

func DemonstrateBlockingCall() {
	var consumer Consumer
	result := make(chan interface{})
	err := make(chan error)
	done := make(chan bool)
	consumer.ExecuteBlockingOp(result, err, done)
waitTillComplete:
	for {
		select {
		case r := <-result:
			fmt.Println("Act on result : ", r)
		case e := <-err:
			fmt.Println("Act on error : ", e)
		case d := <-done:
			if d {
				fmt.Println("Blocking call finished executing.")
				break waitTillComplete
			}
		}
	}
	fmt.Println("Caller continues after execute..")
}
