package progress

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Bar struct {
	Current int64     // The current progress
	Total   int64     // The total we are trying to get to
	Output  io.Writer // Where to output the progress bar
	Width   int       // The character width of the progress bar
}

// Create a new progress Bar. total is the total progress we are trying to make
func NewBar(total int64) *Bar {
	return &Bar{
		Total:  total,
		Output: os.Stdout,
		Width:  40,
	}
}

func (p *Bar) printBar() {
	// Calculate our progress
	percent := 0
	if p.Current > p.Total {
		percent = 100
	} else {
		percent = int((float64(p.Current) / float64(p.Total)) * 100)
	}
	progressLen := 0
	if percent > 0 {
		progressLen = p.Width * percent / 100
	}
	fillLen := p.Width - progressLen

	b := new(bytes.Buffer)
	fmt.Fprintf(b, "\r[%s", strings.Repeat("=", progressLen))
	fmt.Fprintf(b, "%s]%d%%", strings.Repeat(" ", fillLen), percent)
	b.WriteTo(p.Output)
}

// Start up the progress bar
func (p *Bar) Start() {
	p.printBar()
}

// Increment the progress bar by n and update the display
func (p *Bar) Inc(n int64) {
	p.Current = p.Current + n
	p.printBar()
}

// Stop the progress bar
func (p *Bar) Stop() {
	p.Output.Write([]byte("\n"))
}
