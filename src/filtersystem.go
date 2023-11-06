package main

import (
	"bufio"
	"runtime"
)

type filtersystemCmd struct{}

func (cmd *filtersystemCmd) Run(global *Globals) error {
	ret := read_system(global.User, global.Group, global.Shadow, runtime.GOOS)
	if ret == nil {
		// iterate over users, leaving out IDs greater than system threshold
		// ie. >500 on bsd, >1000 on linux
		// then write to files group, master.passwd or shadow + passwd
	}
	return ret
}
