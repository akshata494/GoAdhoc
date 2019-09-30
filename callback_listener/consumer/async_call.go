package main

import (
	"fmt"
	"sync"
)

//DemonstrateAsyncCall - subscribes to events from consumer, then calls consumer's execute
//and parallelly continues executing it's own next steps.
//Currently, waitgroup only adds the execute routine. The functions listening to results are
//mock versions and the waitgroup doesn't add them.
func DemonstrateAsyncCall(wg *sync.WaitGroup) {
	var consumer Consumer
	consumer.Init()
	consumer.Subscribe(waitOnResult, waitOnError, waitOnCompletion)
	consumer.ExecuteAsyncOp(wg)
	fmt.Println("Caller continuing ..")
	wg.Wait()
}

//Dummy result listner.
func waitOnResult(result <-chan interface{}) {
	for {
		select {
		case r := <-result:
			fmt.Println("Acting on result : ", r)
			return
		default:
			//do nothing
		}
	}
}

//Dummy result listner.
func waitOnError(err <-chan error) {
	for {
		select {
		case e := <-err:
			fmt.Println("Acting on error : ", e)
			return
		default:
			//do nothing
		}
	}
}

//Dummy result listner.
func waitOnCompletion(done <-chan bool) {
	for {
		select {
		case c := <-done:
			fmt.Println("Acting on completion : ", c)
			return
		default:
			//do nothing
		}
	}
}
