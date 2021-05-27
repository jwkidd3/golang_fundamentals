package mypkg

var foo = "This is not exported."
var Bar = "This variable IS exported."

func fooFunction(x int) int {
	return x - 1
}

func BarFunction(x int) int {
	return x + 1
}
