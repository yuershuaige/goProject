package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	t.Logf("counter = %d", counter)
}

// Mutex用于保证线程安全 类似于java中的lock
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			mut.Lock()
			counter++
			defer func() {
				mut.Unlock()
			}()
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

// WaitGroup是用于控制流程的 不是用来保证线程安全的 类似于java中的countDownLatch
func TestCounterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mux sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			mux.Lock()
			counter++
			defer func() {
				wg.Done()
				mux.Unlock()
			}()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
