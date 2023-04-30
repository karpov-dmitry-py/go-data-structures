package service

import "fmt"

type LinkedListService interface {
	GetLength() int
	GetValues() []int64
	InsertLeft(value int64)
	InsertRight(value int64)
	GetFirst() (int64, error)
	GetLast() (int64, error)
}

type linkedListService struct {
	list linkedList
}

var (
	emptyListError = fmt.Errorf("list is empty")
)

// NewLinkedListService creates a LinkedList implementation
func NewLinkedListService() LinkedListService {
	return &linkedListService{list: linkedList{nil, 0}}
}

func (l *linkedListService) GetLength() int {
	return l.list.length
}

func (l *linkedListService) GetValues() []int64 {
	var (
		result []int64
	)

	if l.GetLength() == 0 {
		return result
	}

	current := l.list.head
	for {
		result = append(result, current.value)
		if current.next == nil {
			break
		}

		current = current.next
	}

	return result
}

func (l *linkedListService) InsertLeft(value int64) {
	tmp1 := &node{value: value, next: nil}

	if l.list.head == nil {
		l.list.head = tmp1
		l.list.length++
		return
	}

	tmp2 := l.list.head
	l.list.head = tmp1
	tmp1.next = tmp2
	l.list.length++
}

func (l *linkedListService) InsertRight(value int64) {
	tmp := &node{value: value, next: nil}

	if l.list.head == nil {
		l.list.head = tmp
		l.list.length++
		return
	}

	current := l.list.head
	for current.next != nil {
		current = current.next
	}

	current.next = tmp
	l.list.length++
}

func (l *linkedListService) GetFirst() (int64, error) {
	if l.GetLength() == 0 {
		return 0, emptyListError
	}

	return l.list.head.value, nil
}

func (l *linkedListService) GetLast() (int64, error) {
	if l.GetLength() == 0 {
		return 0, emptyListError
	}

	tmp := l.list.head
	for tmp.next != nil {
		tmp = tmp.next
	}

	return tmp.value, nil
}
