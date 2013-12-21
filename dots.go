package clp

import (
	"bytes"
	"io"
	"os"
	"time"
)

type Dots struct {
	Interval       time.Duration // Time until new dot is displayed
	Heads          []byte        // The head of the dots display (the spinning wheel)
	Output         io.Writer     // Where to output the dots
	i              int           // Current index of Heads
	stop, finished chan bool
}

// Create new clp Dots
func NewDots() *Dots {
	return &Dots{
		Interval: time.Second,
		Heads:    []byte{'|', '/', '-', '|', '\\'},
		Output:   os.Stdout,
	}
}

// Start the clp dots
func (p *Dots) Start() {
	p.stop, p.finished = make(chan bool), make(chan bool)
	go func() {
		progress := time.NewTicker(time.Duration(p.Interval))
		update := time.NewTicker(time.Duration(time.Second / 10))
		defer func(){
			progress.Stop()
			update.Stop()
		}()
		b := new(bytes.Buffer)
		for {
			b.Reset()
			select {
			case <-p.stop:
				io.WriteString(p.Output, "\bdone\n")
				p.finished <- true
				close(p.stop)
				return
			case <-progress.C:
				b.WriteString("\b.")
				b.WriteByte(p.head())
				b.WriteTo(p.Output)
			case <-update.C:
				b.WriteByte('\b')
				b.WriteByte(p.head())
				b.WriteTo(p.Output)
			}
		}
	}()
}

// Get the head and advance our index
func (p *Dots) head() byte {
	if p.i >= len(p.Heads) {
		// Reset our index
		p.i = 0
	}
	b := p.Heads[p.i]

	// Advance our index
	p.i = p.i + 1

	return b
}

// Stop the clp dots
func (p *Dots) Stop() {
	p.stop <- true
	<-p.finished
	close(p.finished)
}
