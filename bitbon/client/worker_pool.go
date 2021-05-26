// Copyright Simcord LLC
// This file is part of the Bitbon System Blockchain Node library.
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

package client

import (
	"sync"
)

// Worker is a blocking operation, that takes a close chan. Worker should exit when close is closed
type Worker func(num int, close <-chan struct{})

//------------------------------------------------------------------------------

// workerWrapper takes a Worker implementation and wraps it within a goroutine
// and channel arrangement. The workerWrapper is responsible for managing the
// lifetime of both the Worker and the goroutine.
type workerWrapper struct {
	worker Worker

	// closeChan can be closed in order to cleanly shutdown this worker.
	closeChan chan struct{}

	// closedChan is closed by the run() goroutine when it exits.
	closedChan chan struct{}
}

func newWorkerWrapper(worker Worker, num int) *workerWrapper {
	w := workerWrapper{
		worker:     worker,
		closeChan:  make(chan struct{}),
		closedChan: make(chan struct{}),
	}

	go w.run(num)

	return &w
}

func (w *workerWrapper) run(num int) {
	defer func() {
		close(w.closedChan)
	}()

	w.worker(num, w.closeChan)
}

func (w *workerWrapper) stop() {
	close(w.closeChan)
}

func (w *workerWrapper) waitStopped() {
	<-w.closedChan
}

//------------------------------------------------------------------------------

// Pool is a struct that manages a collection of workers, each with their own
// goroutine.
type Pool struct {
	constructor Worker

	workerMut sync.Mutex
	workers   []*workerWrapper
}

// New creates a new Pool of workers. You must
// provide a constructor function that creates new Worker types and when you
// change the size of the pool the constructor will be called to create each new
// Worker.
func NewPool(constructor Worker) *Pool {
	p := &Pool{
		constructor: constructor,
	}

	return p
}

// Process runs the pool
func (p *Pool) Process(n int) {
	p.SetSize(n)
}

// SetSize changes the total number of workers in the Pool. This can be called
// by any goroutine at any time unless the Pool has been stopped, in which case
// a panic will occur.
func (p *Pool) SetSize(n int) {
	p.workerMut.Lock()
	defer p.workerMut.Unlock()

	lWorkers := len(p.workers)
	if lWorkers == n {
		return
	}

	// Add extra workers if N > len(workers)
	for i := lWorkers; i < n; i++ {
		p.workers = append(p.workers, newWorkerWrapper(p.constructor, i+1))
	}

	// Asynchronously stop all workers > N
	for i := n; i < lWorkers; i++ {
		p.workers[i].stop()
	}

	// Synchronously wait for all workers > N to stop
	for i := n; i < lWorkers; i++ {
		p.workers[i].waitStopped()
	}

	// Remove stopped workers from slice
	p.workers = p.workers[:n]
}

// GetSize returns the current size of the pool.
func (p *Pool) GetSize() int {
	p.workerMut.Lock()
	defer p.workerMut.Unlock()

	return len(p.workers)
}

// Close will terminate all workers and close the job channel of this Pool.
func (p *Pool) Close() {
	p.SetSize(0)
}

//------------------------------------------------------------------------------
