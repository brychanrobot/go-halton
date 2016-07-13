package main

import (
	"fmt"

	"github.com/brychanrobot/halton"
	"github.com/chrisport/go-stopwatch/stopwatch"
)

func main() {
	hSampler := halton.NewHaltonSampler(19)
	sw := stopwatch.NewStopwatch()
	for hSampler.CurrentIndex < 100000000 {
		//fmt.Println(getValue(i, 19))
		value := hSampler.Next()
		if hSampler.CurrentIndex%1000000 == 0 {
			//fmt.Printf("%f, %d ns\n", value, sw.GetPreciseAndRestart())
			sw.LogAndRestart(fmt.Sprintf("%f at index %d", value, hSampler.CurrentIndex))
		}
	}

	fmt.Println("")
}
