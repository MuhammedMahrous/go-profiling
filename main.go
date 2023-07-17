package main

import (
	"flag"
	"github.com/pkg/profile"
	"go-profiling/handler"
	"net/http"
	"runtime"
)

func main() {
	var version = flag.Int("v", 1, "selects the implementation version")
	flag.Parse()

	configureProfiling()
	listenAndServe(*version)
}

func configureProfiling() {
	// Increasing the sampling rate
	runtime.SetCPUProfileRate(1000)

	// Comment out the needed profile
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.GoroutineProfile).Stop()
	//defer profile.Start(profile.BlockProfile).Stop()
	//defer profile.Start(profile.MemProfileHeap).Stop()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
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
