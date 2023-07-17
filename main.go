package main

import (
	"flag"
	"net/http"
	"runtime"

	"github.com/pkg/profile"
	"go-profiling/handler"
)

func main() {
	var version = flag.Int("v", 1, "selects the implementation version")
	flag.Parse()

	// Configure profiling

	// Increasing the sampling rate
	runtime.SetCPUProfileRate(1000)

	// Comment out the needed profile
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.GoroutineProfile).Stop()
	//defer profile.Start(profile.BlockProfile).Stop()
	//defer profile.Start(profile.MemProfileHeap).Stop()
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	listenAndServe(*version)
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
