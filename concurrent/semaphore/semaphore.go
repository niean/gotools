package semaphore

import ()

type Semaphore struct {
	bufSize int
	channel chan int8
}

func NewSemaphore(concurrencyNum int) *Semaphore {
	return &Semaphore{channel: make(chan int8, concurrencyNum), bufSize: concurrencyNum}
}

func (this *Semaphore) Acquire() {
	this.channel <- int8(0)
}

func (this *Semaphore) Release() {
	<-this.channel
}

// not concurrent safe
func (this *Semaphore) AvailablePermits() int {
	return this.bufSize - len(this.channel)
}
