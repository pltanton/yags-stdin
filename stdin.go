package main

import (
	"bufio"
	"os"

	"github.com/spf13/viper"

	"github.com/pltanton/yags/plugins"
)

func (s stdin) StartMonitor() {
	for s.scanner.Scan() {
		s.out <- s.scanner.Text()
	}
}

type stdin struct {
	scanner *bufio.Scanner
	out     chan string
}

func (s stdin) Chan() chan string {
	return s.out
}

func New(conf *viper.Viper) plugins.Plugin {
	return stdin{
		scanner: bufio.NewScanner(os.Stdin),
		out:     make(chan string, 1),
	}
}
