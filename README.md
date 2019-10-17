# GoTasks
A repo of idiomatic code for random ad-hoc tasks in my effort to practice ,learn and share concepts in Go.

Each task is organized in a separate folder. The tasks I'm trying to code for are explained in as much detail as possible below.

Task 1 (callback_listener)
--------------------------
The task is to implement a custom listenable future response struct in Go. These are response types where the operation takes more time to complete and there are more than one result objects are present as part of the response. In that situation, the caller can choose to wait till it is completed (or) caller can resume the next operation with subscribing the events using the response object in a non-blocking way.

The interface Response class has methods like addResult, setError, setCompleted, subscribe(...)

The caller decides not to wait, but interested in listening to the events happening within the response.
Response response = someOperation.execute();
response.subscribe(functional hook to listen to new results, functional hook to listen for error, functional hook for completion)
//exit the current thread as subscribing is non-blocking.


Task 2 (task_queue)
-------------------
The task is to implement a task queue in Go and have a go-routine which periodically checks the tasks queue and inspect if the task is completed or not. If the task is completed then remove it from the queue, if not completed push back into the queue. If the task is not completed after a certain amount of time then it should be removed from the queue and marked as a timeout. 

type Task struct {
   Id string
   IsCompleted boolean // have a random function to mark the IsCompleted after a random period  
   Status string //completed, failed, timeout
}
