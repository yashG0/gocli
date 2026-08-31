package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No command provided")
		return
	}
	argument := os.Args[1]

	switch argument {
	case "info":
		info()
	case "files":
		files()
	case "search":
		search()
	default:
		handleDefault(argument)
	}
}

func info() {
	fmt.Println("System Information")
	fmt.Println("------------------")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Hostname: None")
	} else {
		fmt.Println("Hostname:", name)
	}
	fmt.Println("CPU Cors:", runtime.NumCPU())
}
func files() {
	fmt.Println("this is files result")
}
func search() {
	fmt.Println("this is search result")
}
func handleDefault(arg string) {
	fmt.Println("Unknown Command:", arg)
}
