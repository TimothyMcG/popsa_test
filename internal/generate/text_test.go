package generate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatTimeAway(t *testing.T) {

	tests := []struct {
		start    time.Time
		end      time.Time
		expected string
	}{
		{start: time.Now().Add(-23 * time.Hour), end: time.Now(), expected: "day"},
		{start: time.Date(2024, 9, 13, 0, 0, 0, 0, time.UTC), end: time.Date(2024, 9, 15, 0, 0, 0, 0, time.UTC), expected: "weekend"},
		{start: time.Date(2024, 9, 10, 0, 0, 0, 0, time.UTC), end: time.Date(2024, 9, 12, 0, 0, 0, 0, time.UTC), expected: "couple of days"},
		{start: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), end: time.Date(2024, 9, 7, 0, 0, 0, 0, time.UTC), expected: "week"},
	}

	for _, v := range tests {
		actual := getTimeContext(v.start, v.end)
		if v.expected != actual {
			t.Errorf("expected: %s, but got: %s", v.expected, actual)
		}
	}
}

func TestGenerateTitlesGeneral(t *testing.T) {
	expected := titleStore["general"]
	actual := generateGeneralTitles()
	assert.Equal(t, expected, actual)
}
