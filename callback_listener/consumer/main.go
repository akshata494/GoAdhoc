package main

import (
	"errors"
	"fmt"
	"response"
	"sync"
)

var wg sync.WaitGroup

//Consumer - User defined struct with a method executing a long operation.
type Consumer struct {
	response.Response
}

//ExecuteBlockingOp - dummy time taking operation in consumer
func (consumer *Consumer) ExecuteBlockingOp(result chan interface{}, err chan error, done chan bool) {
	go func() {
		//some long operation that can possibly send multiple results.
		result <- "dummy operation result"
		//if any error
		err <- errors.New("dummy operation error")
		//all done
		done <- true
	}()
}

//IsResultNil - Helper func for AddResult of Response to evaluate if consumer's result is valid.
//Consumer struct needs to have a particular type for the results it produces. Assuming string here.
func (consumer *Consumer) IsResultNil(data interface{}) bool {
	if _, ok := data.(string); ok && data != "" {
		return true
	}
	return false
}

//ExecuteAsyncOp - dummy time taking operation in consumer
func (consumer *Consumer) ExecuteAsyncOp(wg *sync.WaitGroup) {
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		//do some work ..
		consumer.AddResult("dummy operation result", consumer.IsResultNil)
		//encountered error
		consumer.SetError(errors.New("dummy operation error"))
		//finished
		consumer.SetCompletion(true)
	}(wg)
}

func main() {

	fmt.Println("Consumer calls blocking execute.")
	//Demonstrates calling an operation and waiting on its results, errors and
	//finally completion before proceeding with next steps.
	DemonstrateBlockingCall()

	fmt.Println("\n Consumer calling asynchronous execute.")
	//Demonstrates a "Consumer" struct using our "Response" framework to listen
	//on operation execution results while not blocking on its completion.
	DemonstrateAsyncCall(&wg)

	wg.Wait()
}
