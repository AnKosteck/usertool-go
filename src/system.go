package main

import (
	"bufio"
	"fmt"
	"runtime"
)

func read_user_file(userpath string, shadow string, isBsd bool) error {
	return nil
}
func read_group_file(grouppath string, isBsd bool) error {
	return nil
}

func read_system(userpath string, grouppath string, shadowpath string, wantedsystem string) error {
	switch(wantedsystem) {
	case "linux":
		if wantedsystem == runtime.GOOS || (userpath != "/etc/passwd" && grouppath != "/etc/group" && shadowpath != "/etc/shadow") {
			read_user_file(userpath, shadowpath, false)
			read_group_file(grouppath, false)
		} else {
			// mayybe raise error here
		}
	case "bsd":
		if wantedsystem == runtime.GOOS || (userpath != "/etc/master.passwd" && grouppath != "/etc/group") {
			read_user_file(userpath, "UNUSED", true)
			read_group_file(grouppath, true)
		} else {
			// maybe raise error here
		}
	default:
		fmt.Println("Unsupported system type", runtime.GOOS, "cannot continue!")
	}
	return nil
}
