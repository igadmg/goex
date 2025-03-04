package pprofex

import (
	"log"
	"os"
	"runtime/pprof"
)

func WriteCPUProfile(fileName string) (func(), error) {
	f, err := os.Create(fileName + ".prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
		return func() {}, err
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
		return func() {}, err
	}

	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}, nil
}
