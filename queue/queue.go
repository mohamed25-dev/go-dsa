package queue

type Queue struct {
	data []int
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}

	front := q.data[0]
	q.data = q.data[1:len(q.data)]

	return front, true
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Peek() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}

	return q.data[0], true
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}
