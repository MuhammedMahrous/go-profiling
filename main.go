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

	profilerStoppers := configureProfiling(cpuProfiling, memProfiling)
	defer stopProfilers(profilerStoppers)

	listenAndServe(*version)
}

func stopProfilers(s []interface{ Stop() }) {
	for i := range s {
		s[i].Stop()
	}
}

func configureProfiling(cpuProfiling, memProfiling *bool) (profilerStoppers []interface{ Stop() }) {
	// Increasing the sampling rate
	runtime.SetCPUProfileRate(1000)

	switch {
	case cpuProfiling != nil && *cpuProfiling:
		profilerStoppers = append(profilerStoppers, profile.Start(profile.CPUProfile, profile.ProfilePath(".")))
	case memProfiling != nil && *memProfiling:
		profilerStoppers = append(profilerStoppers, profile.Start(profile.MemProfile, profile.ProfilePath(".")))
	}

	return
}

func readArguments() (version *int, cpuProfiling, memProfiling *bool) {
	version = flag.Int("v", 1, "selects the implementation version")
	cpuProfiling = flag.Bool("pcpu", false, "enables CPU profiling")
	memProfiling = flag.Bool("pmem", false, "enables memory profiling")

	flag.Parse()
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
