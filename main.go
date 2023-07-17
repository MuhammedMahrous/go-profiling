package main

import (
	"flag"
	"net/http"
	"runtime"

	"github.com/pkg/profile"

	"go-profiling/handler"
)

func main() {
	version, cpuProfiling, memProfiling := readArguments()

	if profilerStopper := configureProfiling(cpuProfiling, memProfiling); profilerStopper != nil {
		defer profilerStopper.Stop()
	}

	listenAndServe(version)
}

func configureProfiling(cpuProfiling, memProfiling bool) (profilerStopper interface{ Stop() }) {
	// Increasing the sampling rate
	runtime.SetCPUProfileRate(1000)

	if cpuProfiling {
		profilerStopper = profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	} else if memProfiling {
		profilerStopper = profile.Start(profile.MemProfile, profile.ProfilePath("."))
	}

	return
}

func readArguments() (version int, cpuProfiling, memProfiling bool) {
	pVersion := flag.Int("v", 1, "selects the implementation version")
	pCpuProfiling := flag.Bool("pcpu", false, "enables CPU profiling")
	pMemProfiling := flag.Bool("pmem", false, "enables memory profiling")
	flag.Parse()

	version = *pVersion
	cpuProfiling = *pCpuProfiling
	memProfiling = *pMemProfiling

	if cpuProfiling && memProfiling {
		panic("Cannot enable multiple profilers")
	}

	return
}

func listenAndServe(version int) {
	h, err := handler.NewHelloHandler(version)
	if err != nil {
		panic(err)
	}

	http.Handle("/hello", h)

	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}

}
