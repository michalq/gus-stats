package limiter

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Limiter interface {
	Wait()
}

type Limit struct {
	counter    int
	limit      int
	mtx        sync.Mutex
	duration   time.Duration
	ticker     *time.Ticker
	lockChan   chan struct{}
	unlockChan chan struct{}
	firstCall  bool
	isDone     bool
}

func NewLimit(limit int, duration time.Duration) *Limit {
	return &Limit{
		counter:    0,
		limit:      limit,
		duration:   duration,
		firstCall:  true,
		isDone:     false,
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

func (l *Limit) wait(ctx context.Context) {
	if len(l.lockChan) > 0 {
		select {
		case <-l.unlockChan:
			// Wait for unlock
			// Reset lock channel
			<-l.lockChan
		case <-ctx.Done():
			l.isDone = true
		}
	}
}

func (l *Limit) Wait(ctx context.Context) {
	if l.isDone {
		return
	}
	l.up()
	l.mtx.Lock()
	l.counter++
	if l.counter >= l.limit {
		l.debug("Locking: limit exceeded <Duration: %s, Limit: %d>\n", l.duration, l.limit)
		l.lockChan <- struct{}{}
	}
	l.mtx.Unlock()
	l.wait(ctx)
}
