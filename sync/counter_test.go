package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times leaves it as 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		want := 3
		if counter.Value() != want {
			t.Errorf("got %d want %d", counter.Value(), want)
		}
	})
}
