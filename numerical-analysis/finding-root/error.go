package main

import (
	"math"
)

// True Error = True Value - Approximate Value
func TrueError(trueValue float64, approximateValue float64) float64 {
	return trueValue - approximateValue
}

// Absolute True Error = |True Error|
func AbsoluteTrueError(trueValue float64, approximateValue float64) float64 {
	return math.Abs(trueValue - approximateValue)
}

// Relative True Error = True Error / True Value
func RelativeTrueError(trueError float64, trueValue float64) float64 {
	return trueError / trueValue
}

// Absolute Relative True Error = |True Error / True Value|
func AbsoluteRelativeTrueError(trueError float64, trueValue float64) float64 {
	return math.Abs(trueError / trueValue)
}

// Approximate Error = PresentApproximation - PreviousApproximation
// Approximate Error is useful when true values are not known or
// very difficult to obtain.
func ApproximateError(presentApprx float64, previousApprx float64) float64 {
	return presentApprx - previousApprx
}

//Defined as the ratio between the approximate error and
// the present approximation.
func RelativeApproximateError(apprxError float64, presentApprx float64) float64 {
	return (apprxError / presentApprx) * 100
}

func AbsoluteRelativeApproximateError(apprxError float64, presentApprx float64) float64 {
	return math.Abs((apprxError / presentApprx) * 100)
}
