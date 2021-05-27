package main

import (
	"fmt"
)

type IPAddr [4]byte // int8

// Stringer method to output an IPAddr in a nice format
func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

var IPmap = map[string]IPAddr{
	"intuit.com": {104, 101, 138, 8},
	"apple.com":  {17, 172, 224, 47},
}

func main() {
	fmt.Println(IPmap["intuit.com"])
	fmt.Println(IPmap["apple.com"])
	fmt.Println(IPmap["vmware.com"])
}
