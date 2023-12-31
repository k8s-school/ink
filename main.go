package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {

	redPtr := flag.Bool("r", false, "Print in red")
	greenPtr := flag.Bool("g", false, "Print in green (default)")
	bluePtr := flag.Bool("b", false, "Print in blue")
	yellowPtr := flag.Bool("y", false, "Print in yellow")

	helpPtr := flag.Bool("h", false, "Display help")

	// Parse command line arguments
	flag.Parse()

	if *helpPtr {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	text := args[0]

	// Choose color
	var printFunc func(a ...interface{}) string
	printFunc = color.New(color.FgGreen).SprintFunc()
	if *redPtr {
		printFunc = color.New(color.FgRed).SprintFunc()
		fmt.Fprintln(os.Stderr, printFunc(text))
		os.Exit(0)
	} else if *greenPtr {
		printFunc = color.New(color.FgGreen).SprintFunc()
	} else if *bluePtr {
		printFunc = color.New(color.FgBlue).SprintFunc()
	} else if *yellowPtr {
		printFunc = color.New(color.FgYellow).SprintFunc()
	}

	// Print colored text
	fmt.Println(printFunc(text))
}
