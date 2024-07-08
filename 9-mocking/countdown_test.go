package main

import (
	"bytes"
	"testing"
	"time"
)

const (
	sleep = "sleep"
	write = "write"
)

// Implements both Sleeper and io.Writer
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		Countdown(buffer, spySleeper)
		got := buffer.String()
		expected := `3
2
1
Go!`

		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperations{}

		Countdown(spySleeperPrinter, spySleeperPrinter)

		wanted := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		// DeepEqual was giving an error even if the two values are correct
		// That's because of the slice capacity
		assertEquality(t, wanted, spySleeperPrinter.Calls)
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

func assertEquality(t testing.TB, want, got []string) {
	t.Helper()

	raiseError := func(want, got any) {
		t.Errorf("wanted calls %v got %v", want, got)
	}

	if len(want) != len(got) {
		raiseError(want, got)
	}

	for i, v := range want {
		if v != got[i] {
			raiseError(want, got)
		}
	}
}
