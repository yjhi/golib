package jutils

import (
	"math/rand"
	"time"
)

var IsInitRandSeed bool

func InitRandSeed() {
	if IsInitRandSeed != true {
		rand.Seed(time.Now().UnixNano())
		IsInitRandSeed = true

	}
}

func NewRandN(max int) func() int {
	InitRandSeed()
	return func() int {
		return rand.Intn(max + 1)
	}
}

func NewRandMN(min int, max int) func() int {
	InitRandSeed()
	return func() int {
		return rand.Intn(max+1) + min
	}
}

func NewRandN32(max int32) func() int32 {
	InitRandSeed()
	return func() int32 {
		return rand.Int31n(max + 1)
	}
}

func NewRandMN32(min int32, max int32) func() int32 {
	InitRandSeed()
	return func() int32 {
		return rand.Int31n(max+1) + min
	}
}

func NewRandN64(max int64) func() int64 {
	InitRandSeed()
	return func() int64 {
		return rand.Int63n(max + 1)
	}
}

func NewRandMN64(min int64, max int64) func() int64 {
	InitRandSeed()
	return func() int64 {
		return rand.Int63n(max+1) + min
	}
}
