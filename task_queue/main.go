package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const TIMEOUT = 10 * time.Second
const TASK_COMPLETED = "Task completed successfully."
const TASK_TIMED_OUT = "Task timed out."
const TASK_FAILED = "Task failed."

var ErrEmptyQueue = errors.New("can not remove from empty queue")

type Task struct {
	Id          string
	IsCompleted bool
	Status      string
	startTime   time.Time // the time when task was enqueued - temporary field to track timeout
}

type TaskNode struct {
	taskItem Task
	next     *TaskNode
}

//Queue : Task queue with no particular limit on number of nodes.
type Queue struct {
	start, end *TaskNode
	length     int
}

func NewTaskQueue() *Queue {
	return &Queue{nil, nil, 0}
}

//Len : Returns the current number of elements in the Q.
func (queue *Queue) Len() int {
	return queue.length
}

//PrintQueue : Helper func to visualize the Q contents.
func (queue *Queue) PrintQueue() {
	if queue.length == 0 {
		fmt.Println("[PQ] Empty queue!")
		return
	}

	temp := queue.start
	for temp != nil {
		fmt.Printf("Id : %s, IsCompleted : %t, status : %s. \n", temp.taskItem.Id, temp.taskItem.IsCompleted, temp.taskItem.Status)
		temp = temp.next
	}
}

//Enqueue : Add an item at the end.
//			input - task to be appended.
//			Returns no value.
func (queue *Queue) Enqueue(task Task) {
	n := &TaskNode{task, nil}

	if queue.length == 0 {
		queue.start = n
		queue.end = n
	} else {
		queue.end.next = n
		queue.end = n
	}

	queue.length++
}

//Dequeue : Remove an item off the front.
//			Returns an error if queue is empty.
func (queue *Queue) Dequeue() (Task, error) {
	var task Task

	if queue.length == 0 {
		return task, ErrEmptyQueue
	}

	n := queue.start
	if queue.length == 1 {
		queue.start = nil
		queue.end = nil
	} else {
		queue.start = queue.start.next
	}

	queue.length--
	return n.taskItem, nil
}

func (queue *Queue) enqueueTasksDummy() {
	for i := 1; i <= 10; i++ {
		queue.Enqueue(Task{strconv.Itoa(i), false, "", time.Now()})
	}
}

func (queue *Queue) markCompletionDummy() {
	if queue.length == 0 {
		fmt.Println("[mCD] Empty queue!")
	}

	temp := queue.start
	for temp != nil {
		temp.taskItem.IsCompleted = true
		//Sleep for random time between 1-5 secs.
		n := rand.Intn(5)
		time.Sleep(time.Duration(n) * time.Second)
		temp = temp.next
	}
}

func main() {
	taskQueue := NewTaskQueue()
	var wg sync.WaitGroup

	//Generate dummy tasks and push them.
	taskQueue.enqueueTasksDummy()

	//Print initial queue.
	taskQueue.PrintQueue()

	wg.Add(1)

	go monitorQueue(taskQueue, &wg)

	//Start marking task completion at random intervals.
	taskQueue.markCompletionDummy()

	wg.Wait()
}

//For simplicity's sake assuming my routine will monitor the Q for 30 secs.
//Once every 3 seconds ,it checks the task statuses and makes necessary changes.

func monitorQueue(queue *Queue, wg *sync.WaitGroup) {
	endMonitor := time.After(30 * time.Second)
	ticker := time.NewTicker(3 * time.Second)

	fmt.Println("Monitoring routine started at :")
	fmt.Println(time.Now().Format(time.RFC3339))

	defer wg.Done()

	for {
		select {
		case <-endMonitor:
			ticker.Stop()
			fmt.Printf("Exit monitoring.")
			return

		//Ideally a worker should pick my task up, execute and mark the execution time somewhere -
		//for the lack of this full implementation, making the below simplistic assumption.
		//Time elapsed = time that has passed since the task was enqueued.

		case t := <-ticker.C:
			fmt.Println("Monitoring Q status at ", t)

			task, err := queue.Dequeue()
			if err != nil {
				//Log any error in dequeuing. Do not exit monitoring.
				fmt.Println("[mQ] Dequeue failed : ", err)
			} else {
				if !task.IsCompleted {
					timeElapsed := time.Now().Sub(task.startTime)
					if timeElapsed >= TIMEOUT {
						task.Status = TASK_TIMED_OUT
						fmt.Printf("Task %s timed out, removing from Q.\n", task.Id)
					} else if timeElapsed < TIMEOUT {
						queue.Enqueue(task)
					}
				} else if task.IsCompleted {
					//Already dequeued. Set status to completed and log.
					task.Status = TASK_COMPLETED
					fmt.Printf("Task %s completed. Removing from Q.\n", task.Id)
				}
			}

			//Print contents of the Q after each monitor activity.
			queue.PrintQueue()
		}
	}
}
