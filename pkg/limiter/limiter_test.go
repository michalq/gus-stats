package limiter_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/michalq/gus-stats/pkg/limiter"
	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	ctx := context.Background()
	limiters := limiter.NewLimiters(
		[]limiter.Definition{
			{Limit: 3, Duration: 1 * time.Second},
			{Limit: 6, Duration: 5 * time.Second},
		},
		true)

	i := 0
	for {
		limiters.Wait(ctx)
		i++
		fmt.Println(i, time.Now())
		if i > 15 {
			ctx.Done()
			return
		}
	}
	assert.Equal(t, true, true)
}
