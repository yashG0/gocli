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
		if len(os.Args) == 2 {
			fmt.Println("Please provide path to search files and folder!")
		}else{
			files(os.Args[2])
		}
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
func files(p string) {
	items, err := os.ReadDir(p)
	if err != nil {
		fmt.Println("Invalid Path:", p)
		return
	}
	for _, item := range items {
		fmt.Println(item.Name())
	}
}
func search() {
	fmt.Println("this is search result")
}
func handleDefault(arg string) {
	fmt.Println("Unknown Command:", arg)
}
