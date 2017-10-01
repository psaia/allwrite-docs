package util

import "time"

// SetInterval calls a function every n.
func SetInterval(action func(), interval time.Duration) func() {
	t := time.NewTicker(interval)
	q := make(chan struct{})
	go func() {
		for {
			select {
			case <-t.C:
				action()
			case <-q:
				t.Stop()
				return
			}
		}
	}()
	return func() {
		close(q)
	}
}
