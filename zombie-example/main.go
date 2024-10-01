package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(time.Minute * 10)
}
