package main

import (
	"strings"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	got := CreateMessage()
	if !strings.Contains(got, "Hello") {
		t.Errorf("CreateMessage() = %s; want contains 'Hello'", got)
	}
}
