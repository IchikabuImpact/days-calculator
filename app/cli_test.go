package main

import (
	"testing"
	"time"
)

func fixedNow() time.Time {
	return time.Date(2025, time.January, 8, 12, 0, 0, 0, time.UTC)
}

func TestRunCLIFlag(t *testing.T) {
	original := now
	now = fixedNow
	defer func() { now = original }()

	got, err := runCLI([]string{"-days", "5"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "2025/01/03" {
		t.Errorf("got %s want %s", got, "2025/01/03")
	}
}

func TestRunCLIPositional(t *testing.T) {
	original := now
	now = fixedNow
	defer func() { now = original }()

	got, err := runCLI([]string{"10"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "2024/12/29" {
		t.Errorf("got %s want %s", got, "2024/12/29")
	}
}

func TestRunCLIError(t *testing.T) {
	if _, err := runCLI([]string{}); err == nil {
		t.Errorf("expected error for missing args")
	}

	if _, err := runCLI([]string{"abc"}); err == nil {
		t.Errorf("expected error for invalid arg")
	}
}
