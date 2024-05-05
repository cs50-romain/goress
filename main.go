package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"text/tabwriter"

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
	
	tabw := tabwriter.NewWriter(os.Stdout, 1, 1, 10, ' ', 0)
	fmt.Fprintln(tabw, "Process:\tPid:\tMemory (Mb):\t")
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
		fmt.Fprintf(tabw, "%s\t%d\t%d\t\n", process.Name(), process.Pid(), process.Memory())
	}
	tabw.Flush()
}
