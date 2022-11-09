package github_actions_usage_calculator

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name          string
		targetMonth   string
		expectedStart time.Time
		expectedEnd   time.Time
	}{
		{
			name:          time.November.String(),
			targetMonth:   "2022-11",
			expectedStart: time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC),
			expectedEnd:   time.Date(2022, 11, 30, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetRange, err := DecideRange(tt.targetMonth)
			if err != nil {
				panic(err)
			}

			if targetRange.Start != tt.expectedStart {
				t.Errorf("expected %v, but got %v", tt.expectedStart, targetRange.Start)
			}
			if targetRange.End != tt.expectedEnd {
				t.Errorf("expected %v, but got %v", tt.expectedEnd, targetRange.End)
			}
		})
	}
}
