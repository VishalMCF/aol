package main

import (
	"fmt"
	"github.com/arriqaaq/aol/aol"
)

func main() {
	opts := aol.DefaultOptions
	log, err := aol.Open("tmp/", opts)
	if err != nil {
		return
	}
	noOfSegments := log.Segments()
	fmt.Println(noOfSegments)
	for i := 1; i <= noOfSegments; i++ {
		j := 0
		for {
			data, err := log.Read(uint64(i), uint64(j))
			if err != nil {
				continue
			}
			fmt.Println(data)
		}
	}
}
