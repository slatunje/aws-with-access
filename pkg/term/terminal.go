// Copyright © 2018 Sylvester La-Tunje. All rights reserved.

package term

import (
	"os"
	"os/user"
	"github.com/fatih/color"
	"fmt"
	"github.com/slatunje/aws-with-access/pkg/utils"
)

// Terminal declares the structure of a terminal
type Terminal struct {
	User *user.User
	CWD  string
}

// NewTerminal creates and returns `*Terminal` instance
// It will set the current user and working directory
func NewTerminal() *Terminal {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	c, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &Terminal{User: u, CWD: c}
}

// CurrentUser returns the current user.
func (t *Terminal) Username() string {
	return t.User.Username
}

// ===

// Process declares the structure of a process
type Process struct {
	Tern *Terminal
	Attr *os.ProcAttr
	Proc *os.Process
}

// NewProcess will transfer stdin, stdout, and stderr to a new
// process. It will also set the current working directory for
// the shell to start in. Empty Env `[]string`, will cause it
// to call `os.Environ()`
func NewProcess(t *Terminal) *Process {
	var pa = os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   t.CWD,
		Env:   nil,
	}
	return &Process{Attr: &pa, Tern: t}
}

// Start will start up a new shell using the login utility
// Note: "login" supplied twice.
// Note: flag `-flp` means "no prompt for password and pass
// through existing environment variables"
func (p *Process) Start() {
	color.Green(">> Starting a new interactive shell\n When you are done. Type `exit` to leave the session.")
	proc, err := os.StartProcess(utils.LoginPath(), []string{"login", "-flp", p.Tern.Username()}, p.Attr)
	if err != nil {
		panic(err)
	}
	p.Proc = proc
}

// Wait waits for the Process to exit, and then returns a
// ProcessState describing its status and an error, if any.
func (p *Process) Wait() {
	if p.Proc == nil {
		return
	}
	state, err := p.Proc.Wait()
	if err != nil {
		panic(err)
	}
	color.Red(fmt.Sprintf("<< Exited shell. You are no longer in interactive shell: %s\n", state.String()))
}
