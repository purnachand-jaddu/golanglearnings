package queue

import "fmt"

type Queue struct {
	values []int
}

type QueueEmptyError struct {
}

func (e *QueueEmptyError) Error() string {
	return "Cannot remove element if queue is empty"
}

func (q *Queue) insert(val int) {
	q.values = append(q.values, val)
}

func (q *Queue) remove() (int, error) {
	if len(q.values) == 0 {
		return -1, &QueueEmptyError{}
	}
	returnValue := q.values[0]
	q.values = q.values[1:len(q.values)]
	return returnValue, nil
}

func TestQueue() {
	queue := Queue{}
	queue.insert(1)
	queue.insert(2)
	queue.insert(3)
	queue.insert(4)

	i, e := queue.remove()
	printTopOfQueue(i, e)
	i, e = queue.remove()
	printTopOfQueue(i, e)
	i, e = queue.remove()
	printTopOfQueue(i, e)
	i, e = queue.remove()
	printTopOfQueue(i, e)
	i, e = queue.remove()
	printTopOfQueue(i, e)
}

func printTopOfQueue(i int, e error) {
	if e != nil {
		fmt.Print(e)
		return
	}
	fmt.Printf("Item removed from queue %d \n", i)
}
