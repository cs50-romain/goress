package process

import (
	//"fmt"
)

type Process struct {
	name	string
	pid	int
	memory	int
}

func (p *Process) Name() string {
	return p.name
}

func (p *Process) Pid() int {
	return p.pid
}

func (p *Process) Usage() int {
	return p.memory
}

func ListProcesses() {

}

func getProcess() {
	
}

func readProcess() {
	
}
