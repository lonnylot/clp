[![Build Status](https://travis-ci.org/lonnylot/cpl.png?branch=master)](https://travis-ci.org/lonnylot/cpl)
[![Codebot](https://codebot.io/badge/github.com/lonnylot/cpl.png)](http://codebot.io/doc/pkg/github.com/lonnylot/cpl "Codebot")

# clp
`clp` is an easy to use interface for showing users progress updates on the CLI.

## Overview
Command line apps are often just quick utilities. However, 
on occassion they process a lot of data and take a long time to run.
While the app knows exactly what is going on, the user is too often
staring at a blank terminal wondering if the app is still running.

This is where `clp` comes in to play. `clp` makes showing your apps
progress easy and quick!

## Installation
Make sure you have a working GO environment. [See the install instructions](http://golang.org/doc/install).

To install `clp`, simply run:

    go get github.com/lonnylot/clp

## Examples

### Bar
Reading large files can take a while. Don't forget about the user!
It's important to let people know how far along you are so they know everything
is working.

```go
/* forgetmenot.go */
package main

import (
	"os"
	"github.com/lonnylot/clp"
)

func main() {
	file, _ := os.Open("largeFile")
	fInfo, _ := file.Stat()

	// Create our clp bar
	clp := clp.NewBar(fInfo.Size())
	// Start the output
	clp.Start()

	b := make([]byte, 1024)
	for {
		n, err := file.Read(b)

		// Increment clp
		clp.Inc(int64(n))

		// Catch our EOL
		if err != nil {
			break
		}
	}

	// Stop our clp bar
	clp.Stop()
}
```

### Dots
Sometimes you've just gotta process. It takes time and you don't know how
much. These long processes can make people wonder if anything is going on at
all. Now you can easily show people things are still moving with `clp` dots.

```go
/* energizer.go */
package main

import (
	"time"
	"github.com/lonnylot/clp"
)

func main() {
	// Create our clp dots
	clp := clp.NewDots()
	// Start the output
	clp.Start()

	t := time.After(time.Duration(time.Second*10))
	<- t // Wait for the timer...

	// Stop our clp dots
	clp.Stop()
}
```
