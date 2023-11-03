package times

import (
	"fmt"
	"time"
)

type TimeProvider interface {
	Now() time.Time
}

// RealTimeProvider 實際時間
type RealTimeProvider struct{}

func (r RealTimeProvider) Now() time.Time {
	fmt.Println("fishRTP/time/ (r RealTimeProvider) Now()")
	return time.Now()
}

// CustomTimeProvider 自定義時間
type CustomTimeProvider struct {
	FixedTime time.Time
}

func (c CustomTimeProvider) Now() time.Time {
	return c.FixedTime
}
