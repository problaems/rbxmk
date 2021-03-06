package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/anaminus/but"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/library"
	"github.com/anaminus/rbxmk/sources"
	lua "github.com/yuin/gopher-lua"
)

// shortenPath transforms the given path so that it is relative to the working
// directory. Returns the original path if that fails.
func shortenPath(filename string) string {
	if wd, err := os.Getwd(); err == nil {
		if abs, err := filepath.Abs(filename); err == nil {
			if r, err := filepath.Rel(wd, abs); err == nil {
				filename = r
			}
		}
	}
	return filename
}

// ParseLuaValue parses a string into a Lua value. Numbers, bools, and nil are
// parsed into their respective types, and any other value is interpreted as a
// string.
func ParseLuaValue(s string) lua.LValue {
	switch s {
	case "true":
		return lua.LTrue
	case "false":
		return lua.LFalse
	case "nil":
		return lua.LNil
	}
	if number, err := strconv.ParseFloat(s, 64); err == nil {
		return lua.LNumber(number)
	}
	return lua.LString(s)
}

// Std contains interfaces to standard file descriptors. It is meant to be used
// in place of os.Std* so that they can work with tests.
type Std struct {
	in  rbxmk.File
	out rbxmk.File
	err rbxmk.File
}

// CommandUsage contains a description of the command.
const CommandUsage = `rbxmk [ FILE ] [ ...VALUE ]

Receives a file to be executed as a Lua script. If "-" is given, then the script
will be read from stdin instead.

Remaining arguments are Lua values to be passed to the file. Numbers, bools, and
nil are parsed into their respective types in Lua, and any other value is
interpreted as a string. Within the script, these arguments can be received from
the ... operator.`

// Main is the entrypoint to the command. init runs after the World envrionment
// is fully initialized and arguments have been pushed, and before the script
// runs.
func Main(args []string, std Std, init func(rbxmk.State)) error {
	// Parse flags.
	flagset := flag.NewFlagSet(args[0], flag.ExitOnError)
	flagset.Usage = func() {
		fmt.Fprintf(flagset.Output(), CommandUsage)
		flagset.PrintDefaults()
	}
	flagset.Parse(args[1:])
	args = flagset.Args()
	if len(args) == 0 {
		flagset.Usage()
		return nil
	}
	file := args[0]
	args = args[1:]

	// Initialize world.
	world := rbxmk.NewWorld(lua.NewState(lua.Options{
		SkipOpenLibs:        true,
		IncludeGoStackTrace: false,
	}))
	for _, f := range formats.All() {
		world.RegisterFormat(f())
	}
	for _, s := range sources.All() {
		world.RegisterSource(s())
	}
	for _, lib := range library.All() {
		if err := world.Open(lib); err != nil {
			but.Fatal(err)
		}
	}

	world.State().SetGlobal("_RBXMK_VERSION", lua.LString(Version))

	// Add script arguments.
	for _, arg := range args {
		world.State().Push(ParseLuaValue(arg))
	}

	if init != nil {
		init(rbxmk.State{World: world, L: world.State()})
	}

	// Run stdin as script.
	if file == "-" {
		return world.DoFileHandle(std.in, len(args))
	}

	// Run file as script.
	filename := shortenPath(filepath.Clean(file))
	return world.DoFile(filename, len(args))
}

func main() {
	but.IfFatal(Main(os.Args, Std{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}, nil))
}
