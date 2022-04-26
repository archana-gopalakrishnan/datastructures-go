package main

import "fmt"

// Queue represents queue that holds a slice
type Queue struct {
	Items []int
}

// Enqueue adds a value to the queue
func (q *Queue) Enqueue(i int) {
	q.Items = append(q.Items, i)
}

// Dequeue removed value from front of the queue - FIFO
func (q *Queue) Dequeue() int {
	toDequeue := q.Items[0]
	q.Items = q.Items[1:]
	return toDequeue
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
