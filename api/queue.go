package main

// Queue for managing queue as an object
type Queue []string

func (q *Queue) push(s string) {
	*q = append(*q, s)
}

func (q *Queue) pop() {
	if len(*q) == 0 {
		*q = nil
	} else {
		*q = (*q)[1:]
	}
}
