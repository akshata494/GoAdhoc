package response

//Response - providing for communicating result, error and completion signals asynchronously.
type Response struct {
	result chan interface{}
	err    chan error
	done   chan bool
}

//Init - create necessary channels for Response struct.
func (response *Response) Init() {
	response.result = make(chan interface{})
	response.err = make(chan error)
	response.done = make(chan bool)
}

//AddResult - data : custom result value.
//			  isResultNil : user defined function to check if his custom result is nil.
//sends any non nil result on the response's result channel.
func (response *Response) AddResult(data interface{}, isResultNil func(interface{}) bool) {
	if !isResultNil(data) {
		response.result <- data
	}
}

//SetError - err : any "error" from operation execution.
//sends the error on response's err channel.
func (response *Response) SetError(err error) {
	if err != nil {
		response.err <- err
	}
}

//setCompletion - done : true indicates operation completion.
//sends the value on response's done channel.
func (response *Response) SetCompletion(done bool) {
	response.done <- done
}

//Subscribe - waitOnxxx : functional hooks one each to listen to and act upon
//			 			  any values from the response instance in question.
//register to listen on values from response's result, error and done channels.
//essentially spawns a routine for each of user's functional hooks.
func (response *Response) Subscribe(waitOnResult func(<-chan interface{}), waitOnError func(<-chan error),
	waitOnCompletion func(<-chan bool)) {
	go waitOnResult(response.result)
	go waitOnError(response.err)
	go waitOnCompletion(response.done)
}
