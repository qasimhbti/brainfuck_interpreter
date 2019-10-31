// Package utils provides common utilities for "main" packages.
package utils

import (
	"log"
	"os"
	"runtime"
)

// InitLog initializes log format.
func InitLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

// LogStart logs application start.
func LogStart(version string, env string) {
	log.Println("Start")
	log.Printf("Version: %s", version)
	log.Printf("Environment: %s", env)
	log.Printf("Go version: %s", runtime.Version())
	log.Printf("Go max procs: %d", runtime.GOMAXPROCS(0))
}
