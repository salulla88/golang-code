//In custom packages, we do not have a main function. That tells the compiler that this is not a executable but a library

package mypackage //name of file and package are the same

func CalculateAverage(xs[] float64) float64 {
	total:=float64(0)
	for _, x:= range xs {
		total += x
	}

	return total/float64(len(xs))
}
