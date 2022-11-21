package limiter

import (
	"fmt"
	"sync"
	"time"
)

type Limit struct {
	counter    int
	limit      int
	mtx        sync.Mutex
	duration   time.Duration
	ticker     *time.Ticker
	lockChan   chan struct{}
	unlockChan chan struct{}
	firstCall  bool
}

func NewLimit(limit int, duration time.Duration) *Limit {
	return &Limit{
		counter:    0,
		limit:      limit,
		duration:   duration,
		firstCall:  true,
		lockChan:   make(chan struct{}, 1),
		unlockChan: make(chan struct{}, 1),
	}
}

func (l *Limit) up() {
	if !l.firstCall {
		return
	}
	l.firstCall = false
	l.ticker = time.NewTicker(l.duration)
	go (func() {
		for {
			<-l.ticker.C
			l.mtx.Lock()
			l.counter = 0
			if len(l.unlockChan) == 0 {
				l.debug("Unlocking, time elapsed <Duration: %s, Limit: %d>\n", l.duration, l.limit)
				l.unlockChan <- struct{}{}
			} else {
				l.debug("Waiting for unlock <Duration: %s, Limit: %d>\n", l.duration, l.limit)
			}
			l.mtx.Unlock()
		}
	})()
}

func (l *Limit) debug(s string, params ...any) {
	if false {
		fmt.Printf(s, params...)
	}
}

func (l *Limit) wait() {
	if len(l.lockChan) > 0 {
		// Wait for unlock
		<-l.unlockChan
		// Reset lock channel
		<-l.lockChan
	}
}

func (l *Limit) Increment() {
	l.up()
	l.mtx.Lock()
	l.counter++
	if l.counter >= l.limit {
		l.debug("Locking: limit exceeded <Duration: %s, Limit: %d>\n", l.duration, l.limit)
		l.lockChan <- struct{}{}
	}
	l.mtx.Unlock()
	l.wait()
}

func Xyz() {
	limit1s := NewLimit(3, 1*time.Second)
	limit15m := NewLimit(10, 30*time.Second)
	// limit12h := NewLimit(1000, 12 * time.Hour)
	// limit7d := NewLimit(10000, 7 * 24 * time.Hour)

	i := 0
	for {
		limit1s.Increment()
		limit15m.Increment()

		i++
		fmt.Println(i, time.Now())
	}
}
