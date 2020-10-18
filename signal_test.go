package gutil

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestWaitCallbackReturnNil(t *testing.T) {
	callback := func(s os.Signal) error {
		return nil
	}

	go func() {
		time.Sleep(time.Second)
		PanicErr(syscall.Kill(syscall.Getpid(), syscall.SIGINT))
	}()
	err := Wait(callback, syscall.SIGTERM, syscall.SIGINT)
	assert.Nil(t, err)
}

func TestWaitCallbackReturnError(t *testing.T) {
	myErr := errors.New("abc")
	callback := func(s os.Signal) error {
		return myErr
	}

	go func() {
		time.Sleep(time.Second)
		PanicErr(syscall.Kill(syscall.Getpid(), syscall.SIGTERM))
	}()
	err := Wait(callback, syscall.SIGTERM)
	assert.Equal(t, myErr, err)
}

func TestWaitUntilGetSignal(t *testing.T) {
	callback := func(s os.Signal) error {
		return nil
	}
	var slept = false
	go func() {
		time.Sleep(3 * time.Second)
		slept = true
		PanicErr(syscall.Kill(syscall.Getpid(), syscall.SIGTERM))
	}()
	_ = Wait(callback, syscall.SIGTERM)
	assert.True(t, slept)
}

func TestWaitForSignalGetSignal(t *testing.T) {
	var slept = false
	go func() {
		time.Sleep(3 * time.Second)
		slept = true
		PanicErr(syscall.Kill(syscall.Getpid(), syscall.SIGTERM))
	}()
	WaitForSignals(syscall.SIGTERM)
	assert.True(t, slept)
}
