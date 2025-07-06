package main

import (
	"testing"
	"time"
)

func TestDateNDaysAgo(t *testing.T) {
	// Save original now function and restore after test
	originalNow := now
	defer func() { now = originalNow }()

	// Fixed time: 2025-01-08 (YYYY-MM-DD). Format uses year-month-day etc.
	fixedTime := time.Date(2025, time.January, 8, 12, 0, 0, 0, time.UTC)
	now = func() time.Time { return fixedTime }

	tests := []struct {
		daysAgo int
		want    string
	}{
		{0, "2025/01/08"},
		{1, "2025/01/07"},
		{7, "2025/01/01"},
	}

	for _, tt := range tests {
		got := dateNDaysAgo(tt.daysAgo)
		if got != tt.want {
			t.Errorf("dateNDaysAgo(%d) = %s, want %s", tt.daysAgo, got, tt.want)
		}
	}
}
