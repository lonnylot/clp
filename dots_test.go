package progress

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func ExampleDots() {
	progress := NewDots()
	progress.Start()
	// ... do something here
	progress.Stop()
}

func TestDotOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	dots := &Dots{
		Interval: time.Second,
		Heads:    []byte{'|', '/', '-', '|', '\\'},
		Output:   buf,
	}

	dots.Start()
	a := time.After(time.Duration(time.Second * 5))
	<-a
	dots.Stop()
	line, err := buf.ReadString('\n')
	if err != nil {
		t.Errorf("Reading buffer string failed with: %s", err.Error())
	}

	if strings.Count(".", line) > 2 {
		t.Errorf("Less than 2 dots were output in 5 seconds")
	}
}

func TestDoneOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	dots := &Dots{
		Interval: time.Second,
		Heads:    []byte{'|', '/', '-', '|', '\\'},
		Output:   buf,
	}

	dots.Start()
	a := time.After(time.Duration(time.Second * 2))
	<-a
	dots.Stop()
	line, err := buf.ReadString('\n')
	if err != nil {
		t.Errorf("Reading buffer string failed with: %s", err.Error())
	}

	if strings.Contains(line, "done") == false {
		t.Errorf("'done' was not output")
	}
}
