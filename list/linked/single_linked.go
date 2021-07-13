package linked

import (
	"fmt"
	"time"

	"github.com/DmitriiTrifonov/algo/list"
)

// SingleLinkedList implementation
type SingleLinkedList struct {
	unixTime int64
	next     *SingleLinkedList
}

// NewSingleLinkedList creates a SingleLinkedList with nil next pointer
func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		next:     nil,
		unixTime: time.Now().UnixNano(),
	}
}

// Add adds a new element to list
func (s *SingleLinkedList) Add() {
	if s.next == nil {
		s.next = NewSingleLinkedList()
		return
	}
	currentNext := s.next
	s.next = NewSingleLinkedList()
	s.next.next = currentNext
}

// Reverse the list
func (s *SingleLinkedList) Reverse() list.List {
	var last *SingleLinkedList
	current := s
	next := current.next
	for {
		last = current
		current = next
		next = next.next
		current.next = last
		if next == nil {
			break
		}
	}
	s.next = nil
	return current

}

// Length counts list elements
func (s *SingleLinkedList) Length() (len int) {
	current := s
	for len = 1; current.next != nil; len++ {
		current = current.next
	}
	return
}

// GetAddresses gets an elements addresses at string format
func (s *SingleLinkedList) GetAddresses() (addr []string) {
	current := s
	length := s.Length()
	addr = make([]string, length)
	for i := range addr {
		addr[i] = fmt.Sprintf("%p", current)
		current = current.next
	}
	return
}

// Slice converts list to slice
func (s *SingleLinkedList) Slice() (sl []list.List) {
	current := s
	sl = make([]list.List, s.Length())
	for i := range sl {
		sl[i] = current
		current = current.next
	}
	return
}
