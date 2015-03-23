package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestMain(T *testing.T) {
	rand.Seed(time.Now().Unix())
}
