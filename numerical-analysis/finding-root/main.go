package main

import (
	"fmt"
	"math"
//	"gonum.org/v1/gonum/diff/fd"
)

type fnc func(x float64) float64

func SecantMethod(f fnc, x00 float64, x01 float64, maxIters int) {
	c := 0
	for c < maxIters {
		x1 := x00 - ( (f(x00) * (x00 - x01)) / (f(x00) - f(x01)))
		ae := ApproximateError(x1, x00)
		arae := AbsoluteRelativeApproximateError(ae, x1)
		fmt.Printf("%d. ||%f || %f || %f || %f\n", c, x01, x00, x1, arae)
		x01 = x00
		x00 = x1
		c++
	}

}

func BisectionMethod(f fnc, xl float64, xu float64, maxIters int, minErr float64) {
	c := 0
	err := 1.0
	previous := -1.0
	xm := (xl + xu/2)
	for c < maxIters && err > minErr{ 
				xm = (xl + xu) / 2


		if f(xl)*f(xm) < 0 {
			xu = xm
		} else if f(xu)*f(xm) < 0 {
			xl = xm
		} else if f(xl)*f(xm) == 0 {
			fmt.Printf("%d. iter Xl: %.4f, xu: %.4f, xm: %.4f, Ea: %.4f, fxm: %.4f\n",
				c+1, xl, xu, xm, err, f(xm))
			return
		}


		err = AbsoluteRelativeApproximateError(ApproximateError(xm, previous), xm)
		previous = xm
		c++
		fmt.Printf("%d. iter Xl: %.4f, xu: %.4f, xm: %.4f, Ea: %.4f, fxm: %.4f\n",
			c, xl, xu, xm, err, f(xm))


	}

}

func ApproximateDerivative(x float64, h float64, f func(x float64) float64) float64 {
	return float64((f(x+h) - f(x)) / h)
}

func f(x float64) float64 {
	return 1 - math.Log(x)
}

/*
func f(x float64) float64 {
	return x*x*x - 0.165*x*x + (3.993 * math.Pow(10, -4))
}

func ff(x float64) float64 {
	return x*x*x - 3*x*x + 5*x + 1
}

func fff (x float64) float64 {
	return math.Pow(1 + math.Log(x), 3)
}
*/
func main() {
	SecantMethod(f, 0.5, 0.1, 4)
	/*
	BisectionMethod(ff, -1, 0, 40, 0.5)
	fmt.Println("*****************\n")
	vals := []float64{0.2, 0.1, 0.01, 0.001}
	previous := 0.0
	for _, val := range vals {
		ad := ApproximateDerivative(2, val, fff)
		trueValue := 4.3001
		ate := AbsoluteTrueError(trueValue, ad)
		arte := AbsoluteRelativeTrueError(ate, trueValue)
		arae := AbsoluteRelativeApproximateError(ApproximateError(ad, previous), ad)
		fmt.Printf("%.4f || %.4f || %.4f || %.4f || %.4f\n", val, ad, ate, arte, arae)
		previous = ad
	} 
	*/

}
