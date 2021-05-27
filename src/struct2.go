package main

import "fmt"

type Mountain struct {
	name      string
	elevation int
}

// START OMIT
func main() {
	everest := Mountain{elevation: 8848, name: "Mt. Everest"}
	fmt.Printf("%#v\n", everest)
	fmt.Println(Mountain{name: "Flat"})
	fmt.Println(&Mountain{name: "Long's Peak", elevation: 4346})
	nd := Mountain{name: "Nanda Devi", elevation: 7815}
	fmt.Println(nd.name) // dot works for structs...
	sp := &nd
	fmt.Println(sp.elevation) // ...and also for pointers to structs
	sp.elevation++
	fmt.Println(sp.elevation)
}

// END OMIT
