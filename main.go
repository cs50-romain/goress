package main

import (
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"os"
)

func main() {
	processes, err := ps.Processes()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("NAME\t|\tPID")
	for _, process := range processes {
		fmt.Printf("%s\t|\t%d\n", process.Executable(), process.Pid())
	}
}
