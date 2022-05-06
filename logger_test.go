package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func stdoutCapExec(fn func()) string {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = orig }()
	fn()
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	out := <-outC
	log.Printf("%T %[1]v", out)
	return out
}

// this is likely to fail all the time because the details of this log
// are very fragile
func TestNewLogger(t *testing.T) {
	exp := "710209223959 [34mINFO[0m dbboilerplate/logger_test.go:47 Hi Tester\n"

	year, month, day, hour, min, sec, nsec, loc :=
		1971, time.February, 9, 22, 39, 59, int(4e8), time.UTC
	mockTime := time.Date(year, month, day, hour, min, sec, nsec, loc)
	patch := monkey.Patch(time.Now, func() time.Time { return mockTime })
	defer patch.Unpatch()
	lf := "/tmp/TestNewLogger_test.log"
	l := NewLogger(lf, "debug")
	defer os.Remove(lf)
	// defer func() {  }()
	l.Infof("Hi %s", "Tester")
	act, err := os.ReadFile(lf)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, exp, string(act))
}
