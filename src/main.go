/**
* Main file
**/
package main

import (
	"fmt"
	"github.com/alecthomas/kong"
)

type Globals struct {
	SystemType string      `help:"Defines the type of system the usertool is execed on. Either b or l"`
	Version    VersionFlag `name:"version" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type TOOLARGS struct {
	Globals

	bsd2linux      bsd2linuxCmd      `cmd: "" help: "Convert BSD-style files to Linux style files."`
	filtersystem   filtersystemCmd   `cmd: "" help: "Filters out system users and prints to stdout."`
	flawsearch     flawsearchCmd     `cmd: "" help: "Takes password/user files and displays flaws inside."`
	linux2bsd      linux2bsdCmd      `cmd: "" help: "Convert Linux style to BSD style"`
	passwordupdate passwordupdateCmd `cmd: "" help: "Update all user passwords from database."`
	removeusers    removeusersCmd    `cmd: "" help: "Remove one or more users from the system."`
	system2yaml    system2yamlCmd    `cmd: "" help: "Convert the system users to a yaml file (Ansible readable)."`
	systemupdate   systemupdateCmd   `cmd: "" help: "Write yaml 'database' to system files."`
	updatedatabase updatedatabaseCmd `cmd: "" help: "Update the yaml database with system changes."`
}

func main() {
	toolargs := TOOLARGS{
		Globals: Globals{
			Version: VersionFlag("0.0.1"),
		},
	}
	kong_context := kong.Parse(&toolargs,
		kong.Name("usertool-go"),
		kong.Description("A simple tool for Unix user management."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			version: "0.0.1",
		},
	)
	// Call the Run() method of the selected parsed command.
	err := kong_context.Run(&toolargs.Globals)
	kong_context.FatalIfErrorf(err)
}
