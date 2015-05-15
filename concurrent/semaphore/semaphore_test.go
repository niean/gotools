package semaphore

import (
	"runtime"
	"sync"
	"testing"
)

func TestSemaphore(t *testing.T) {
	sema := NewSemaphore(2)
	sema.Acquire()
	sema.Acquire()

	if sema.AvailablePermits() != 0 {
		t.Error("AvailablePermits, AvailablePermits")
	}

	sema.Release()
	if sema.AvailablePermits() != 1 {
		t.Error("AvailablePermits, AvailablePermits")
	}
	sema.Release()
}

func BenchmarkSemaphore(b *testing.B) {
	b.StopTimer()
	sema := NewSemaphore(1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sema.Acquire()
		sema.Release()
	}
}

func BenchmarkSemaphoreConcurrent(b *testing.B) {
	b.StopTimer()
	sema := NewSemaphore(1)
	wg := sync.WaitGroup{}
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for i := 0; i < each; i++ {
				sema.Acquire()
				sema.Release()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
