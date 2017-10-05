/* gobook pg 123
4. You may have noticed that every function in the
packages we've seen start with a capital letter.
In Go if something starts with a capital letter
that means other packages (and programs) are
able to see it. If we had named the function
average instead of Average our main program
would not have been able to see it.
*/
package math

import m "math"

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return 0
	}
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Find the minimum of a series of numbers
func Min(xs []float64) float64 {
	min := m.MaxFloat64
	if len(xs) == 0 {
		return 0
	}
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

// Find the maximum of a series of numbers
func Max(xs []float64) float64 {
	max := -m.MaxFloat64
	if len(xs) == 0 {
		return 0
	}
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}
