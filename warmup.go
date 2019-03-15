package warmup

import (
	"context"
	"time"
)

// LimitFunc will receive the number seconds that have passed since the limiter was created and should
// return the amount of tokens we should give to the bucket per second.
type LimitFunc func(x float64) float64

// Limiter provides a way to increase a rate limit over the course of time. You can specify a custom
// function to determine how many tokens the bucket should be refilled with.
type Limiter struct {
	f LimitFunc
	start time.Time
	throttle chan time.Time
}

// NewLimiter will create a new limiter with a given function, and the bucket size.
func NewLimiter(f LimitFunc, b int) *Limiter {
	l := Limiter{
		f:        f,
		throttle: make(chan time.Time, b),
		start: time.Now(),
	}

	// This will start the ticker, that will increase in frequency over time
	go l.tick()

	return &l
}

// WaitN will wait until N tokens are ready to be used before returning
func (l *Limiter) WaitN(ctx context.Context, n int) error {
	for i := 0; i < n; i++ {
		<-l.throttle
	}

	return nil
}

// tick will publish the current time on the channel at every tick. Tick will use the given
// LimitFunc to calculate how long it should wait between ticks.
func (l *Limiter) tick() <-chan time.Time {
	for {
		x := time.Since(l.start).Seconds()
		y := l.f(x)

		// We need to check to make sure y is not less than 1
		if y < 1 {
			y = 1
		}

		// Wait the determined amount of time
		rate := time.Second / time.Duration(int64(y))
		time.Sleep(rate)

		l.throttle <- time.Now()
	}
}
