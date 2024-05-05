package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	name	string
	pid	int
	memory	int // in Mb
}

func (p *Process) Name() string {
	return p.name
}

func (p *Process) Pid() int {
	return p.pid
}

func (p *Process) Memory() int {
	return p.memory
}

func ListProcesses() {

}

func GetProcess(pid int) (*Process, error) {
	processfile := fmt.Sprintf("/proc/%d/stat", pid)
	file, err := os.Open(processfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var processInfo string
	buf := make([]byte, 1024)
	for  {
		str, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if str > 0 {
			processInfo += string(buf[:str])
		}
	}

	processInfoSeparated := strings.Split(processInfo, " ")

	/*
	memoryInfo, err := ReadMemoryStats()
	if err != nil {
		return nil, err
	}
	*/
	
	processMemoryUsageConverted, err := strconv.Atoi(processInfoSeparated[22])
	if err != nil {
		return nil, err
	}

	processPid, err := strconv.Atoi(processInfoSeparated[0])
	if err != nil {
		return nil, err
	}
	processName := strings.Trim(processInfoSeparated[1], "(")
	processName = strings.Trim(processName, ")")

	var processMemoryUsage int
	processMemoryUsage = processMemoryUsageConverted / 1000000
	
	// Calculate memory usage percentage
	/*
	if processMemoryUsageConverted == 0 {
		processMemoryUsage = 0
	} else {
		processMemoryUsage = memoryInfo.MemTotal / (processMemoryUsageConverted * 1000)
	}
	*/

	return &Process{name: processName, pid: processPid, memory: processMemoryUsage,}, nil
}

type memory struct {
	MemTotal int
	MemFree  int
}

func ReadMemoryStats() (*memory, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	mem := &memory{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		memoryStats := parseLine(scanner.Text())
		memoryKey := memoryStats[0]
		memoryValue, err := strconv.Atoi(memoryStats[1])
		if err != nil {
			//fmt.Printf("memory: %s\n", err)
			continue
		}
		if memoryKey == "MemTotal" {
			mem.MemTotal = memoryValue
		} else if memoryKey == "MemFree" {
			mem.MemFree = memoryValue
		}
	}
	return mem, nil
}

func parseLine(raw string) []string {
	raw = strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	return strings.Split(raw, ":")
}
