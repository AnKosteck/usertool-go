package main

import (
	"bufio"
)

type linux2bsdCmd struct{}

func (cmd *linux2bsdCmd) Run(global *Globals) error {
	// read out the files
	ret := read_system(global.User, global.Group, global.Shadow, "linux")
	if ret == nil {
		// filter out system users + groups
		// write to group, master.passwd
		return nil
    }
	return ret
}
