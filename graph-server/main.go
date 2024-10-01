package main

import (
	"fmt"
	"github.com/alecthomas/kingpin/v2"
	"net/http"
	"os"
	"path/filepath"
)

var (
	app      = kingpin.New("graph-server", "Graph server application to render graphs.")
	dataPath = app.Flag("data-path", "Path to json data files").Default("/data").String()
	check    = app.Command("check-legacy-format", "Checks legacy json formats.")
	migrate  = app.Command("migrate-legacy", "Migrates legacy format files.")
	run      = app.Command("run", "Runs a web server to serve graphs.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case migrate.FullCommand():
		fmt.Fprintln(os.Stdout, "Migrating graphs...")
	case check.FullCommand():
		fmt.Fprintf(os.Stdout, "Checking %s directory for legacy format...\n", *dataPath)

		filepath.WalkDir(*dataPath, func(path string, d os.DirEntry, err error) error {
			if !d.IsDir() {
				fmt.Fprintf(os.Stderr, "File %s has legacy format\n", path)
			}

			return nil
		})
	case run.FullCommand():
		fmt.Fprintln(os.Stdout, "Starting http server on :8080...")
		http.HandleFunc("/healthcheck", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "hello\n")
		})
		http.ListenAndServe(":8080", nil)
	}
}
