package lighthouse

import (
	"fmt"
	"math"
)

func dashOffset(score int) string {
	// Circumference = 2 * pi * 54 = 339.292
	circumference := 2 * math.Pi * 54
	offset := circumference * (1 - float64(score)/100)
	return fmt.Sprintf("%.3fpx", offset)
}

func intToStr(n int) string {
	return fmt.Sprintf("%d", n)
}
