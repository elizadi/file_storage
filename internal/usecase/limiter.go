package usecase

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"sync"
)

var (
	MaxViewErr = errors.New("Already maximum view connection. Try later")
)

type Limiter struct {
	counter int
	mutex   sync.Mutex
	limit   int
}

func (v *Limiter) Acquire() error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	if v.counter >= v.limit {
		logger.Error(MaxViewErr)
		return MaxViewErr
	}

	v.counter++
	return nil
}

func (v *Limiter) Release() {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	if v.counter <= 0 {
		panic("Release without acquire")
	}
	v.counter = v.counter - 1
}

func NewLimiter(limit int) Limiter {
	if limit <= 0 {
		panic(errors.New("limit must be greater than 0"))
	}
	return Limiter{
		counter: 0,
		mutex:   sync.Mutex{},
		limit:   limit,
	}
}
