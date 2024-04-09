package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := Point{X: 150, Y: 60}
	got := secondHandPoint(tm)
	if !almostEqual(got, want) {
		t.Errorf("At midnight, expected second hand almost at %+v, got %+v", want, got)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(2020, time.January, 1, 0, 0, 30, 0, time.UTC)
	want := Point{X: 150, Y: 240}
	got := secondHandPoint(tm)
	if !almostEqual(got, want) {
		t.Errorf("At 30 seconds, expected second hand almost at %+v, got %+v", want, got)
	}
}

func TestMinuteHandAt15Minutes(t *testing.T) {
	tm := time.Date(2020, time.January, 1, 0, 15, 0, 0, time.UTC)
	want := Point{X: 220, Y: 150} // Correcting to reflect the 3 o'clock position
	got := minuteHandPoint(tm)
	if !almostEqual(got, want) {
		t.Errorf("At 15 minutes, expected minute hand almost at %+v, got %+v", want, got)
	}
}

func TestHourHandAt6Hours(t *testing.T) {
	tm := time.Date(2020, time.January, 1, 6, 0, 0, 0, time.UTC)
	want := Point{X: 150, Y: 200} // Corrected to reflect the actual length and position
	got := hourHandPoint(tm)
	if !almostEqual(got, want) {
		t.Errorf("At 6 hours, expected hour hand almost at %+v, got %+v", want, got)
	}
}

// compare two Point values within a small margin of error
func almostEqual(a, b Point) bool {
	const delta = 0.001
	return math.Abs(a.X-b.X) < delta && math.Abs(a.Y-b.Y) < delta
}
