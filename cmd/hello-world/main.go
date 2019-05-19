package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/powerman/structlog"
)

//nolint:gochecknoglobals
var (
	cmd = strings.TrimSuffix(path.Base(os.Args[0]), ".test")
	ver string // set by ./build
	log = structlog.New()
	cfg struct {
		version  bool
		logLevel string
	}
)

// Init provides common initialization for both app and tests.
func Init() {
	flag.BoolVar(&cfg.version, "version", false, "print version")
	flag.StringVar(&cfg.logLevel, "log.level", "debug", "log `level` (debug|info|warn|err)")

	log.SetDefaultKeyvals(
		structlog.KeyUnit, "main",
	)
}

func main() {
	Init()
	flag.Parse()

	switch {
	case cfg.version: // Must be checked after all other flags for ease testing.
		fmt.Println(cmd, ver, runtime.Version())
		os.Exit(0)
	}

	// Wrong log.level is not fatal, it will be reported and set to "debug".
	structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(cfg.logLevel))
	log.Info("started", "version", ver)

	http.HandleFunc("/", greet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
	log.Info("greeting")
}
