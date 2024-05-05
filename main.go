package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	ps "cs50-romain/goress/process"
)

func main() {
	nums := regexp.MustCompile("[0-9]+")
	// Read /proc directory and find any dir that starts with a num
	dirs, err := os.ReadDir("/proc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	for _, dir := range dirs {
		if !nums.Match([]byte(dir.Name())) {
			continue
		}

		pid, err := strconv.Atoi(dir.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		process, err := ps.GetProcess(pid)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("process: %s, pid: %d, memory: %dMb\n", process.Name(), process.Pid(), process.Memory())
	}
}
