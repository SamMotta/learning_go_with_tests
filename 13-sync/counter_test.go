package main

import (
	"sync"
	"testing"
)

func Test(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCounter := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCounter)

		for range wantedCounter {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCounter)
	})

}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
