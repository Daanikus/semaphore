package semaphore_test

import (
	"os"
	"testing"
	"time"

	"github.com/kamilsk/semaphore"
)

func TestMultiplex(t *testing.T) {
	sleep := 100 * time.Millisecond

	start := time.Now()
	<-semaphore.Multiplex(semaphore.WithSignal(os.Interrupt), semaphore.WithTimeout(sleep))
	end := time.Now()

	if expected, obtained := sleep, end.Sub(start); expected > obtained {
		t.Errorf("an unexpected sleep time. expected: %v; obtained: %v", expected, obtained)
	}
}

func TestMultiplex_WithoutChannels(t *testing.T) {
	<-semaphore.Multiplex()
}

func TestWithDeadline(t *testing.T) {
	tests := []struct {
		name     string
		deadline time.Duration
	}{
		{"normal case", 10 * time.Millisecond},
		{"past deadline", -time.Nanosecond},
	}
	for _, test := range tests {
		start := time.Now()
		<-semaphore.WithDeadline(start.Add(test.deadline))
		end := time.Now()

		if !end.After(start.Add(test.deadline)) {
			t.Errorf("%s: an unexpected deadline", test.name)
		}
	}
}

func TestWithSignal_NilSignal(t *testing.T) {
	<-semaphore.WithSignal(nil)
}

func TestWithTimeout(t *testing.T) {
	sleep := 100 * time.Millisecond

	start := time.Now()
	<-semaphore.WithTimeout(sleep)
	end := time.Now()

	if expected, obtained := sleep, end.Sub(start); expected > obtained {
		t.Errorf("an unexpected sleep time. expected: %v; obtained: %v", expected, obtained)
	}
}
