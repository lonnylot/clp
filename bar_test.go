package clp

import (
	"bytes"
	"strings"
	"testing"
)

func ExampleBar() {
	clp := NewBar(1024)
	clp.Start()
	for i := 0; i < 2; i++ {
		clp.Inc(512)
	}
	clp.Stop()
}

func TestWidth(t *testing.T) {
	buf := new(bytes.Buffer)
	bar := &Bar{
		Total:  1024,
		Width:  10,
		Output: buf,
	}
	bar.Start()
	bar.Inc(1024)
	bar.Stop()

	line, err := buf.ReadString('\n')
	if err != nil {
		t.Errorf("Reading buffer string failed with: %s", err.Error())
	}

	if strings.Contains(line, "==========") == false {
		t.Errorf("Bar width was not long enough.")
	}
}

func TestPercentage(t *testing.T) {
	buf := new(bytes.Buffer)
	bar := &Bar{
		Total:  1024,
		Width:  10,
		Output: buf,
	}
	bar.Start()
	bar.Inc(512)
	bar.Stop()

	line, err := buf.ReadString('\n')
	if err != nil {
		t.Errorf("Reading buffer string failed with: %s", err.Error())
	}

	if strings.Contains(line, "50%") == false {
		t.Errorf("Bar percent was incorrect. %s", line)
	}

	buf.Reset()
	bar.Start()
	bar.Inc(512)
	bar.Stop()

	line, err = buf.ReadString('\n')
	if err != nil {
		t.Errorf("Reading buffer string failed with: %s", err.Error())
	}

	if strings.Contains(line, "100%") == false {
		t.Errorf("Bar percent was incorrect. %s", line)
	}
}
