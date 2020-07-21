package library

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/anaminus/rbxmk"
	"github.com/yuin/gopher-lua"
)

func OS(s rbxmk.State) {
	lib, ok := s.L.GetGlobal("os").(*lua.LTable)
	if !ok {
		lib = s.L.CreateTable(0, 5)
		s.L.SetGlobal("os", lib)
	}
	lib.RawSetString("split", s.WrapFunc(osSplit))
	lib.RawSetString("join", s.WrapFunc(osJoin))
	lib.RawSetString("expand", s.WrapFunc(osExpand))
	lib.RawSetString("getenv", s.WrapFunc(osGetenv))
	lib.RawSetString("dir", s.WrapFunc(osDir))
}

func osSplit(s rbxmk.State) int {
	path := s.L.CheckString(1)
	for i := 2; i <= s.L.GetTop(); i++ {
		var result string
		switch typ := s.L.CheckString(i); typ {
		case "dir":
			result = filepath.Dir(path)
		case "base":
			result = filepath.Base(path)
		case "ext":
			result = filepath.Ext(path)
		case "stem":
			result = filepath.Base(path)
			result = result[:len(result)-len(filepath.Ext(path))]
		case "fext":
			// result = scheme.GuessFileExtension(ctx.Options, "", path)
			// if result != "" && result != "." {
			// 	result = "." + result
			// }
		case "fstem":
			// ext := scheme.GuessFileExtension(ctx.Options, "", path)
			// if ext != "" && ext != "." {
			// 	ext = "." + ext
			// }
			// result = filepath.Base(path)
			// result = result[:len(result)-len(ext)]
		default:
			s.L.RaiseError("unknown argument %q", typ)
			return 0
		}
		s.L.Push(lua.LString(result))
	}
	return s.L.GetTop() - 1
}

func osJoin(s rbxmk.State) int {
	j := make([]string, s.L.GetTop())
	for i := 1; i <= s.L.GetTop(); i++ {
		j[i-1] = s.L.CheckString(i)
	}
	filename := filepath.Join(j...)
	s.L.Push(lua.LString(filename))
	return 1
}

func osExpand(s rbxmk.State) int {
	expanded := os.Expand(s.L.CheckString(1), func(v string) string {
		switch v {
		case "script_name", "sn":
			if fi, ok := s.PeekFile(); ok {
				path, _ := filepath.Abs(fi.Path)
				return filepath.Base(path)
			}
		case "script_directory", "script_dir", "sd":
			if fi, ok := s.PeekFile(); ok {
				path, _ := filepath.Abs(fi.Path)
				return filepath.Dir(path)
			}
		case "working_directory", "working_dir", "wd":
			wd, _ := os.Getwd()
			return wd
		case "temp_directory", "temp_dir", "tmp":
			return os.TempDir()
		}
		return ""
	})
	s.L.Push(lua.LString(expanded))
	return 1
}

func osGetenv(s rbxmk.State) int {
	value, ok := os.LookupEnv(s.L.CheckString(1))
	if ok {
		s.L.Push(lua.LString(value))
	} else {
		s.L.Push(lua.LNil)
	}
	return 1
}

func osDir(s rbxmk.State) int {
	dirname := s.L.CheckString(1)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		s.L.RaiseError(err.Error())
		return 0
	}
	tfiles := s.L.CreateTable(len(files), 0)
	for _, info := range files {
		tinfo := s.L.CreateTable(0, 4)
		tinfo.RawSetString("name", lua.LString(info.Name()))
		tinfo.RawSetString("isdir", lua.LBool(info.IsDir()))
		tinfo.RawSetString("size", lua.LNumber(info.Size()))
		tinfo.RawSetString("modtime", lua.LNumber(info.ModTime().Unix()))
		tfiles.Append(tinfo)
	}
	s.L.Push(tfiles)
	return 1
}