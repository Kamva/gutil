package gutil

import (
	"os"
	"os/signal"
)

// Wait waits until get one signal of your specified signals, then
// calls to the callback and return result of callback as result of the
// Wait function.
func Wait(callback func(s os.Signal) error, signals ...os.Signal) (err error) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, signals...)

	go func() {
		sig := <-sigs
		err = callback(sig)
		done <- true
	}()

	<-done
	return
}
