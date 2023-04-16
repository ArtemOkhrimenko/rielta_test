package app

import (
	"sync/atomic"
)

var (
	Number = int64(0)
)

func (a *App) GetNumber() int64 {

	return Number
}

func (a *App) PutNumber(num int64) int64 {
	newNum := atomic.AddInt64(&Number, num)

	return newNum
}
