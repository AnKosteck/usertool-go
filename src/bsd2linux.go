package main

import (
	"bufio"
)

type bsd2linuxCmd struct{}

func (cmd *bsd2linuxCmd) Run(global *Globals) error {
	// read out system
	ret := read_system(global.MasterPasswd, global.Group, "UNUSED", "bsd")
	if ret == nil {
		// filter out system users
		// write to passwd, group, shadow files
		return nil
	}
	return ret
}
