package main

import (
	"flag"
	"log"
	"os/exec"
	"time"
)

var (
	startCoordinator = flag.Bool("start-coordinator", false, "start coordinator")
	startWorker      = flag.Bool("start-worker", false, "start worker")
)

func startLongRunningImageOptimizationWorker() {
	cmd := exec.Command("sleep", "30")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

func startImageOptimizationCoordinator() {
	cmd := exec.Command("/image-optimization", "--start-worker")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	if *startCoordinator {
		log.Println("starting a coordinator...")
		startImageOptimizationCoordinator()
		time.Sleep(time.Minute * 10)
	}

	if *startWorker {
		startLongRunningImageOptimizationWorker()
		time.Sleep(time.Second * 10)
	}
}
