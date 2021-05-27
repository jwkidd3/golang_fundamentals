package main

// START OMIT
type ByteSize float64

const (
	_           = iota // ignore first value by assigning to empty
	KB ByteSize = 1 << (10 * iota) // << is left shift operator
	MB
	GB
	TB
	PB
	EB
	ZB
)

func main() {
	println("GB =", GB)
	println("TB =", TB)
}
//END OMIT
