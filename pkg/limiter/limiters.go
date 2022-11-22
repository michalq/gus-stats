package limiter

import (
	"context"
	"time"
)

type Definition struct {
	Limit    int
	Duration time.Duration
}

type Limiters []Limiter

func (lms Limiters) Wait(ctx context.Context) {
	for _, limiter := range lms {
		limiter.Wait(ctx)
	}
}

func NewLimiters(definitions []Definition, debugMode bool) Limiters {
	lms := make(Limiters, 0, len(definitions))
	for _, definition := range definitions {
		lms = append(lms, newLimit(definition.Limit, definition.Duration, debugMode))
	}
	return lms
}
