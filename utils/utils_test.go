package utils

import (
	"bytes"
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	InitLog()
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	LogStart("1.0.0", "testing")
}
