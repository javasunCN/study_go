package day4

import (
	"sync"
	"testing"
)

/**
Mutex:Mutex 和信道来处理竞态条件，Go协程锁
	Lock()
	unLock()
如果在任意时刻只允许一个 Go 协程访问临界区，那么就可以避免竞态条件。而使用 Mutex 可以达到这个目的。

Mutex 用于提供一种加锁机制（Locking Mechanism），可确保在某时刻只有一个协程在临界区运行，以防止出现竞态条件。

Mutex vs 信道:
	1.Go 协程需要与其他协程通信时，可以使用信道
	2.当只允许一个协程访问临界区时，可以使用 Mutex

*/
// 实例一:含有竞态条件的程序
var x = 0

func Test_Mutex_1(log *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	log.Logf("主协程结束 %d", x)
}
func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

// 实例二:处理竞态 Mutex
func Test_Mutex_2(log *testing.T) {

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementMux(&wg, &mu)
	}
	wg.Wait()
	log.Logf("主协程结束 %d", x)
}
func incrementMux(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	// Mutex.Lock 与 Mutex.Unlock 之间只允许一个协程操作
	x = x + 1
	mu.Unlock()
	wg.Done()
}

// 实例二:使用信道处理竞态
func Test_Mutex_3(log *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementChan(&wg, ch)
	}
	wg.Wait()
	log.Logf("主协程结束 %d", x)
}
func incrementChan(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
