package pron

import (
	"bytes"
	"fmt"
)

// Interface for external and internal jobs
type job interface {
	Register(schedule string, tab *Prontab) error
	Dispatch() ([]byte, error)
}

// External job rollup w/ time maps
type externalJob struct {
	s      schedule
	action externalAction
}

// Internal job rollup w/ time maps
type internalJob struct {
	s      schedule
	action internalAction
}

// Registers an external job to the tab
func (a *externalJob) register(p *Prontab) {
	p.j.externalJobs = append(p.j.externalJobs, *a)
}

// Registers an internal job to the tab
func (a *internalJob) register(p *Prontab) {
	p.j.internalJobs = append(p.j.internalJobs, *a)
}

// Internal action dispatch
func (j *internalJob) Dispatch() ([]byte, error) {
	return j.action.fn()
}

// TODO: fix + cleanup console routines
// External action dispatch
func (j *externalJob) Dispatch() ([]byte, error) {
	var buf bytes.Buffer

	cmd := j.action.cmd
	cmd.Stdout = &buf

	err := cmd.Start()
	out := buf.Bytes()

	fmt.Println(buf.String())
	buf.Reset()
	cmd.Process.Kill()

	return out, err

	// return j.action.cmd.Output()
}

func (j *externalJob) scheduled(t tick) bool {
	return scheduled(t, j.s)
}

func (j *internalJob) scheduled(t tick) bool {
	return scheduled(t, j.s)
}
