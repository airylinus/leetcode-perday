package main

import (
	"container/list"
	"fmt"
)

// LRUContainer contains jobs
type LRUContainer struct {
	LastJob *Job
	// Jobs    [5]*Job
	KeyMap map[string]*Job
}

// Job ..
type Job struct {
	Name string
	Next *Job
}

func main() {
	l := list.New()
	l.PushFront(1)
	fmt.Print(l)
	// lc := &LRUContainer{}
	// job1 := Job{Name: "job1"}
	// lc.Put(&job1)
	// lc.Put(&Job{Name: "job2"})
	// lc.Put(&Job{Name: "job3"})
	// lc.Put(&Job{Name: "job4"})
	// lc.Put(&Job{Name: "job5"})
	// lc.Put(&Job{Name: "job6"})
}

// func (lc *LRUContainer)

// Get returns job instance
func (lc *LRUContainer) Get(name string) *Job {
	tmp := lc.LastJob
	for tmp.Next != nil {
		if tmp.Name == name {
			return tmp
		}
		tmp = tmp.Next
	}
	return nil
}

// Put puts job instance
func (lc *LRUContainer) Put(job *Job) {
	old := lc.LastJob
	if _, ok := lc.KeyMap[job.Name]; !ok {
		lc.LastJob = job
		lc.KeyMap[job.Name] = job
	}
	if old != nil {
		job.Next = old
	}
	// max size is 5
	tmp := lc.LastJob
	for i := 0; i < 5; i++ {
		if i == 4 {
			tmp.Next = nil
		}
		tmp = tmp.Next
	}
}
