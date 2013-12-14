# progress
`progress` is an easy to use API for keeping users up to date on the CLI.

## Overview
Command line apps are often just quick utilities. However, 
on occassion they process a lot of data and take a long time to run.
While the app knows exactly what is going on, the user is too often
staring at a static terminal wondering if the app is still running.

This is where `progress` comes in to play. `progress` makes it showing your apps
status easy and quick!

## Installation
Make sure you have a working GO environment. [See the install instructions](http://golang.org/doc/install).

To install `progress`, simply run:

    go get github.com/lonnylot/progress

## Examples

### Bar
Reading large files can take a while. Don't forget about the user!
It's important to let people know where you are so they know everything is 
working.

```go
/* forgetmenot.go */
package main

import (
	"os"
	"github.com/lonnylot/progress"
)

func main() {
	file, _ := os.Open("largeFile")
	fInfo, _ := file.Stat()

	// Create our progress bar
	progress := progress.NewBar(fInfo.Size())
	// Start the output
	progress.Start()

	b := make([]byte, 1024)
	for {
		n, err := file.Read(b)

		// Increment progress
		progress.Inc(int64(n))

		// Catch our EOL
		if err != nil {
			break
		}
	}

	// Stop our progess bar
	progress.Stop()
}
```

### Dots
Sometimes you've just gotta process. It takes time and you don't know how
much. These long processes can make people wonder if anything is going on at
all. Now you can easily show people things are still moving.

```go
/* energizer.go */
package main

import (
	"time"
	"github.com/lonnylot/progress"
)

func main() {
	// Create our progress dots
	progress := progress.NewDots()
	// Start the output
	progress.Start()

	t := time.After(time.Duration(time.Second*10))
	<- t // Wait for the timer...

	// Stop our progress dots
	progress.Stop()
}
```
