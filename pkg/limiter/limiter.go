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
	select {
	case <-l.unlockChan:
		l.debug("Release %s", l.duration)
		// Do nothing
	case <-ctx.Done():
		if len(l.unlockChan) > 0 {
			<-l.unlockChan
		}
		l.isDone = true
	}
}

func (l *Limit) Wait(ctx context.Context) {
	if l.isDone {
		return
	}
	l.up()
	l.mtx.Lock()
	l.counter++
	isLimitExceeded := l.counter >= l.limit
	l.mtx.Unlock()
	if isLimitExceeded {
		l.debug("Locking: limit exceeded <Duration: %s, Limit: %d>", l.duration, l.limit)
		l.wait(ctx)
	}
}
