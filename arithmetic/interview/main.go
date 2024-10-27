package main

import "sync"

func main() {
	var lock1 sync.Mutex
	var lock2 sync.Mutex

	lock1.Lock()
	lock2.Lock() // 死锁发生在这里，因为lock1已经被锁定

	lock1.Unlock()
	lock2.Unlock()
}
