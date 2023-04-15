package repository

import "time"

type TimeRepository interface {
	TimeNow() time.Time
}

type TimeRepositoryCtx struct{}

func (c *TimeRepositoryCtx) TimeNow() time.Time {
	return time.Now()
}
