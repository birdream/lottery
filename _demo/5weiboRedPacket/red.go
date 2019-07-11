package main

import (
	"math/rand"
	"sync"
	"time"
)

type task struct {
	id       uint32
	callback chan uint
}

// var packageList map[uint32][]uint = make(map[uint32][]uint)
var packageList *sync.Map = new(sync.Map)
var chTasks chan task = make(chan task)

func fetchPLMoney() {
	for {
		t := <-chTasks
		id := t.id
		l, ok := packageList.Load(id)

		if ok && l != nil {
			list := l.([]uint)

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			i := r.Intn(len(list))
			money := list[i]

			if len(list) > 1 {
				if i == len(list)-1 {
					// packageList[uint32(id)] = list[:i]
					packageList.Store(uint32(id), list[:i])
				} else if i == 0 {
					// packageList[uint32(id)] = list[1:i]
					packageList.Store(uint32(id), list[1:i])
				} else {
					// packageList[uint32(id)] = append(list[:i], list[i+1:]...)
					packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
				}
			} else {
				// delete(packageList, uint32(id))
				packageList.Delete(uint32(id))
			}

			t.callback <- money
		} else {
			t.callback <- 0
		}

	}
}
