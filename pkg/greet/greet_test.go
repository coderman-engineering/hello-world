package greet_test

import (
	"github.com/coderman-engineering/hello-world/pkg/greet"
	"testing"
)

func TestDefaultGreeting(t *testing.T) {
	expected := greet.GreetingHelloWorld
	got := greet.DefaultGreeting()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
