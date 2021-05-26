// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
// This file has been modified, you can find the original version by following the link
// <https://github.com/ethereum/go-ethereum>
//
// The Bitbon System Blockchain Node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Bitbon System Blockchain Node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Bitbon System Blockchain Node library. If not, see <http://www.gnu.org/licenses/>.

// Package prque implements a priority queue data structure supporting arbitrary
// value types and int64 priorities.
//
// If you would like to use a min-priority queue, simply negate the priorities.
//
// Internally the queue is based on the standard heap package working on a
// sortable version of the block based stack.
package prque

import (
	"container/heap"
)

// Priority queue data structure.
type Prque struct {
	cont *sstack
}

// New creates a new priority queue.
func New(setIndex SetIndexCallback) *Prque {
	return &Prque{newSstack(setIndex)}
}

// Pushes a value with a given priority into the queue, expanding if necessary.
func (p *Prque) Push(data interface{}, priority int64) {
	heap.Push(p.cont, &item{data, priority})
}

// Peek returns the value with the greates priority but does not pop it off.
func (p *Prque) Peek() (interface{}, int64) { // nolint:gocritic
	item := p.cont.blocks[0][0]
	return item.value, item.priority
}

// Pops the value with the greates priority off the stack and returns it.
// Currently no shrinking is done.
func (p *Prque) Pop() (interface{}, int64) { // nolint:gocritic
	item := heap.Pop(p.cont).(*item)
	return item.value, item.priority
}

// Pops only the item from the queue, dropping the associated priority value.
func (p *Prque) PopItem() interface{} {
	return heap.Pop(p.cont).(*item).value
}

// Remove removes the element with the given index.
func (p *Prque) Remove(i int) interface{} {
	if i < 0 {
		return nil
	}
	return heap.Remove(p.cont, i)
}

// Checks whether the priority queue is empty.
func (p *Prque) Empty() bool {
	return p.cont.Len() == 0
}

// Returns the number of element in the priority queue.
func (p *Prque) Size() int {
	return p.cont.Len()
}

// Clears the contents of the priority queue.
func (p *Prque) Reset() {
	*p = *New(p.cont.setIndex)
}
