package main

import (
	"fmt"
	"os"
	"path/filepath"
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
		if len(os.Args) < 3 {
			fmt.Println("Please provide path to search files and folder!")
		} else {
			files(os.Args[2])
		}
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a filename!")
		} else if len(os.Args) == 3 {
			total := search(os.Args[2], ".")
			fmt.Println("Total:",total,"found!")
		} else {
			total:= search(os.Args[2], os.Args[3])
			fmt.Println("Total:",total,"found!")
		}
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
	fmt.Println("CPU Cores:", runtime.NumCPU())
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

func search(filename string, p string) int {
	items, err := os.ReadDir(p)
	count := 0
	if err != nil {
		fmt.Println("Invalid Path:", p)
		return 0
	}
	for _, item := range items {
		if item.IsDir() {
			count += search(filename, filepath.Join(p, item.Name()))
		} else {
			if item.Name() == filename {
				fmt.Println("Found:", filepath.Join(p, item.Name()))
				count++
			}
		}
	}
	return count
}
func handleDefault(arg string) {
	fmt.Println("Unknown Command:", arg)
}
