package main

import (
	"flag"
	"inputoutput"
	"log"
	"mandelbrot"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)

		defer pprof.StopCPUProfile()
	}

	inputoutput.Init()

	mandelbrot.LinkScreenOutput(inputoutput.ScreenOutputChan)
	mandelbrot.LinkInput(inputoutput.InputChan)

	go mandelbrot.DrawMandelbrot()

	inputoutput.Run()
}