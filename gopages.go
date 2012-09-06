// Copyright 2010 Abiola Ibrahim <abiola89@gmail.com>. All rights reserved.
// Use of this source code is governed by New BSD License
// http://www.opensource.org/licenses/bsd-license.php
// The content and logo is governed by Creative Commons Attribution 3.0
// The mascott is a property of Go governed by Creative Commons Attribution 3.0
// http://creativecommons.org/licenses/by/3.0/

package main

import (
	"code.google.com/p/gopages/util"
	"errors"
	"flag"
	"os"
	"os/exec"
)

const (
	MAKE    = iota
	GOBUILD = iota
)

//where the build execution starts
func main() {
	cl := flag.Bool("clean", false, "don't build, just clean the generated pages")
	//run := flag.Bool("run", false, "run the generated executable after build")
	flag.Parse()
	if *cl {
		err := clean()
		if err != nil {
			println(err.Error())
		}
		return
	}
	settings, err := util.LoadSettings() //inits the settings and generates the .go source files
	if err != nil {
		println(err.Error())
		return
	}
	util.Config = settings.Data //stores settings to accessible variable
	println("generated", len(settings.Data["pages"]), "gopages")
	err = util.AddHandlers(settings.Data["pages"]) //add all handlers
	if err != nil {
		println(err.Error())
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "get" {
		build()
	} else {
		println("run \"gopages get\" to build with go get after generating pages")
	}
}

//create the pages directory to store generated source codes
func init() {
	err := os.MkdirAll(util.DIR, 0755)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

//to build the project with gobuild or make after generating .go source files
func build() (err error) {
	fd := []*os.File{os.Stdin, os.Stdout, os.Stderr}
	goexec, _ := exec.LookPath("go")
	if len(goexec) == 0 {
		return errors.New("go not found in PATH")
	}
	dir := os.Getenv("PWD")
	//	id, err := os.ForkExec(gofmt, []string{"", "-w", file}, os.Environ(), dir, fd)
	process, err := os.StartProcess(goexec, []string{"", "get"}, &os.ProcAttr{Env: os.Environ(), Files: fd, Dir: dir})
	if err != nil {
		return
	} else {
		process.Wait()
		//os.Wait(id, 0)
	}
	return
}

//deletes the generated source codes
func clean() (err error) {
	err = os.RemoveAll(util.DIR)
	if err != nil {
		println(err.Error())
	}
	return
}
