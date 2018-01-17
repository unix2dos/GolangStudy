package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func main() {
	a := &Once{}
	a.Do(func() { fmt.Println("22") })
	a.Do(func() { fmt.Println("22") })
}
