package usecase

import "testing"

func TestDoGenerateCounter(t *testing.T) {
	t.Run("", func(t *testing.T) {
		want := "09"
		got := doGenerateCounter(8)
		if got != want {
			t.Errorf("I want %s but I got %s\n", want, got)
		}
	})
}
