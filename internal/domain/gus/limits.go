package gus

import (
	"time"

	"github.com/michalq/gus-stats/pkg/limiter"
)

func BaseApiLimits(debug bool) limiter.Limiters {
	return limiter.NewLimiters([]limiter.Definition{
		{Limit: 5, Duration: 2 * time.Second}, // 1s but sometimes there are errors so lets stick with 2s
		{Limit: 100, Duration: 15 * time.Minute},
		{Limit: 1000, Duration: 12 * time.Hour},
		{Limit: 10000, Duration: 7 * 24 * time.Hour},
	}, debug)
}

func RegisteredApiLimits(debug bool) limiter.Limiters {
	return limiter.NewLimiters([]limiter.Definition{
		{Limit: 10, Duration: 2 * time.Second}, // 1s but sometimes there are errors so lets stick with 2s
		{Limit: 500, Duration: 15 * time.Minute},
		{Limit: 5000, Duration: 12 * time.Hour},
		{Limit: 50000, Duration: 7 * 24 * time.Hour},
	}, debug)
}
