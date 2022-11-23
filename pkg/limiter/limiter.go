package limiter

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Limiter interface {
	Wait(context.Context)
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
	debugMode  bool
}

func newLimit(limit int, duration time.Duration, debug bool) *Limit {
	return &Limit{
		counter:    0,
		limit:      limit,
		duration:   duration,
		firstCall:  true,
		isDone:     false,
		debugMode:  debug,
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
				l.debug("Unlocking, time elapsed <Duration: %s, Limit: %d>", l.duration, l.limit)
				l.unlockChan <- struct{}{}
			} else {
				l.debug("Waiting for unlock <Duration: %s, Limit: %d>", l.duration, l.limit)
			}
			l.mtx.Unlock()
		}
	})()
}

func (l *Limit) debug(s string, params ...any) {
	if l.debugMode {
		fmt.Printf("[%s] %s\n", time.Now(), fmt.Sprintf(s, params...))
	}
}

func (l *Limit) wait(ctx context.Context) {
	// TODO Deadlock hazard between checking len() and reseting? To verify.
	if len(l.lockChan) > 0 {
		select {
		case <-l.unlockChan:
			// Reset lock channel
			<-l.lockChan
		case <-ctx.Done():
			// Release locks, set to done.
			<-l.lockChan
			if len(l.unlockChan) > 0 {
				<-l.unlockChan
			}
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
		l.debug("Locking: limit exceeded <Duration: %s, Limit: %d>", l.duration, l.limit)
		l.lockChan <- struct{}{}
	}
	l.mtx.Unlock()
	l.wait(ctx)
}
