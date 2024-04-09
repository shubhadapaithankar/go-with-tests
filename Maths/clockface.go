package clockface

import (
	"fmt"
	"math"
	"time"
)

const (
	clockCenterX     = 150
	clockCenterY     = 150
	secondHandLength = 90
	minuteHandLength = 70
	hourHandLength   = 50
)

// A Point represents a point in 2D space.
type Point struct {
	X float64
	Y float64
}

// calculates the end point of the second hand.
func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t), secondHandLength)
}

// calculates the end point of the minute hand.
func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t), minuteHandLength)
}

// calculates the end point of the hour hand.
func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t), hourHandLength)
}

// converts current seconds into angle in radians.
func secondsInRadians(t time.Time) float64 {
	return math.Pi / 30 * float64(t.Second())
}

// converts current minute into angle in radians.
func minutesInRadians(t time.Time) float64 {
	return (math.Pi / 30 * float64(t.Minute())) + (math.Pi / 1800 * float64(t.Second()))
}

// converts current hour into angle in radians.
func hoursInRadians(t time.Time) float64 {
	return (math.Pi / 6 * float64(t.Hour()%12)) + (math.Pi / 360 * float64(t.Minute()))
}

// converts an angle in radians to a point on the clock face.
func angleToPoint(angle float64, length float64) Point {
	x := clockCenterX + math.Sin(angle)*length
	y := clockCenterY - math.Cos(angle)*length
	return Point{X: x, Y: y}
}

// generates the SVG representation of the clock face.
func GenerateSVG(t time.Time) string {
	return fmt.Sprintf(`<svg width="300" height="300" xmlns="http://www.w3.org/2000/svg">
<circle cx="%[1]d" cy="%[2]d" r="100" style="fill:white; stroke:black; stroke-width:4"/>
<line x1="%[1]d" y1="%[2]d" x2="%[3]f" y2="%[4]f" style="stroke:black;stroke-width:2"/>
<line x1="%[1]d" y1="%[2]d" x2="%[5]f" y2="%[6]f" style="stroke:black;stroke-width:2"/>
<line x1="%[1]d" y1="%[2]d" x2="%[7]f" y2="%[8]f" style="stroke:red;stroke-width:2"/>
</svg>`,
		clockCenterX, clockCenterY,
		hourHandPoint(t).X, hourHandPoint(t).Y,
		minuteHandPoint(t).X, minuteHandPoint(t).Y,
		secondHandPoint(t).X, secondHandPoint(t).Y)
}
