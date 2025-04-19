package main

import (
	"fmt"
	"os"
	"radu.macocian.me/goboy/cpu"
	"radu.macocian.me/goboy/errorHandling"
	"radu.macocian.me/goboy/memory"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing rom, please provide a rom")
		return
	}

	checkExtension(os.Args[1])

	data, err := os.ReadFile(os.Args[1])
	check(err)

	startMem := uint(1000)
	memory.WriteAll(startMem, data)
	cpu.Execute(startMem)
}

func checkExtension(file string) {
	splits := strings.Split(file, ".")
	if splits[len(splits)-1] != "gb" {
		panic(errorHandling.InvalidRomError)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
