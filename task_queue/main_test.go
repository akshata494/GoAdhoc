package main

import (
	"strconv"
	"testing"
	"time"
)

//helper func for deep equal on two queues excluding the time field
func compareQueues(x, y *Queue) bool {
	if x.Len() == 0 && y.Len() == 0 {
		return true
	} else if x.Len() != y.Len() {
		return false
	}
	tx := x.start
	ty := y.start

	for tx != nil || ty != nil {
		xtask := tx.taskItem
		ytask := ty.taskItem
		if xtask.Id != ytask.Id || xtask.Status != ytask.Status ||
			xtask.IsCompleted != ytask.IsCompleted {
			return false
		}
		tx = tx.next
		ty = ty.next
	}
	return true
}

func compareTasks(xtask, ytask Task) bool {
	if xtask.Id != ytask.Id || xtask.Status != ytask.Status ||
		xtask.IsCompleted != ytask.IsCompleted {
		return false
	}
	return true
}

func TestEnqueue(t *testing.T) {
	//Case : Enqueue should append the Task item to end of queue

	now := time.Now()
	testQueue := NewTaskQueue()
	inputTask := Task{strconv.Itoa(10), false, "", now}
	wantQueue := &Queue{
		start:  &TaskNode{Task{strconv.Itoa(10), false, "", now}, nil},
		end:    nil,
		length: 1,
	}

	testQueue.Enqueue(inputTask)
	if !compareQueues(testQueue, wantQueue) {
		t.Errorf("Enqueue - want %+v, got %+v", wantQueue, testQueue)
	}
}

func TestDequeue(t *testing.T) {
	//Case 1 : Dequeue on an empty queue should return error "can not remove from empty queue".
	//Case 2 : Dequeue should remove and return the first item in the queue.

	emptyQueue := Queue{
		start:  nil,
		end:    nil,
		length: 0,
	}
	inputQueue := Queue{
		start:  &TaskNode{Task{strconv.Itoa(20), false, "", time.Now()}, nil},
		end:    nil,
		length: 1,
	}
	wantTask := &TaskNode{Task{strconv.Itoa(20), false, "", time.Now()}, nil}
	cases := []struct {
		input interface{}
		want  interface{}
	}{
		{emptyQueue, ErrEmptyQueue},
		{inputQueue, wantTask},
	}

	_, err := emptyQueue.Dequeue()
	if err != cases[0].want {
		t.Errorf("emptyQueue.Dequeue() : want %s, got %s", cases[0].want, err)
	}

	got, _ := inputQueue.Dequeue()
	if !compareTasks(got, wantTask.taskItem) {
		t.Errorf("Dequeue - want %+v, got %+v", cases[1].want, got)
	}
}
