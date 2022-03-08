package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

// 方式三
//var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// 方式二
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	// 方式三
	//cmdLine.StringVar(&name, "name", "everyone", "The greeting object")
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {

	//方式一
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	// 方式三
	//_ = cmdLine.Parse(os.Args[1:])

	flag.Parse()
	fmt.Printf("Hello, %s!\n\n", name)
}
