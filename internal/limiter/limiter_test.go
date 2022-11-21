package limiter_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/michalq/gus-stats/internal/limiter"
	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	ctx := context.Background()
	limit1s := limiter.NewLimit(3, 1*time.Second)
	limit15m := limiter.NewLimit(6, 5*time.Second)

	i := 0
	for {
		limit1s.Wait(ctx)
		limit15m.Wait(ctx)

		i++
		fmt.Println(i, time.Now())
		if i > 15 {
			ctx.Done()
			return
		}
	}
	assert.Equal(t, true, true)
}
