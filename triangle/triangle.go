// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "Not A Triangle" // not a triangle
	Equ Kind = "equilateral"    // equilateral
	Iso Kind = "isosceles"      // isosceles
	Sca Kind = "scalene"        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	// default not a triangle
	k = NaT
	// filter out NaNs and infinities, I don't know if this is the right thing to do here
	if (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)) || (math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0)) {
		return k
	}
	// Is it really a triangle?
	if (a > 0 && b > 0 && c > 0) && (a+b >= c && a+c >= b && c+b >= a) {
		if a == b && b == c {
			k = Equ
		} else if a == b || a == c || c == b {
			k = Iso
		} else {
			k = Sca
		}
	}

	return k
}
