The code for this assignment is organized as below :
	|____ response (package response)
		|_____response.go 
			|______ Response struct 
			|______ Init()
			|______ AddResult()
			|______ SetError ()
			|______ SetCompletion ()
			|______ Subscribe()
	|_____ consumer (package main)
		|_____ main.go 
			|_____ Consumer struct (consuming Response struct for illustration)
			|_____ ExecuteBlockingOp()
			|_____ ExecuteAsyncOp()
		|_____ blocking_call.go 
			|_____ DemonstrateBlockingCall ()
		|_____ async_call.go
			|_____ DemonstrateAsyncCall()
			|_____ waitOnResult() 
			|_____ waitOnError()
			|_____ waitOnCompletion()
			
			
Build :
With the above code in $GOPATH/src, 
cd response
go install 
cd ../
cd consumer 
go run main.go async_call.go blocking_call.go