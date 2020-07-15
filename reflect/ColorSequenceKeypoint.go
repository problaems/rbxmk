package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func ColorSequenceKeypoint() Type {
	return Type{
		Name:        "ColorSequenceKeypoint",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(types.ColorSequenceKeypoint).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "ColorSequenceKeypoint").(types.ColorSequenceKeypoint)
				return s.Push("bool", v.(types.ColorSequenceKeypoint) == op)
			},
		},
		Members: map[string]Member{
			"Time": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.ColorSequenceKeypoint).Time)
			}},
			"Value": {Get: func(s State, v Value) int {
				return s.Push("Color3", v.(types.ColorSequenceKeypoint).Value)
			}},
			"Envelope": {Get: func(s State, v Value) int {
				return s.Push("float", v.(types.ColorSequenceKeypoint).Envelope)
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.ColorSequenceKeypoint
				switch s.Count() {
				case 2:
					v.Time = s.Pull(1, "float").(float32)
					v.Value = s.Pull(2, "Color3").(types.Color3)
				case 3:
					v.Time = s.Pull(1, "float").(float32)
					v.Value = s.Pull(2, "Color3").(types.Color3)
					v.Envelope = s.Pull(3, "float").(float32)
				default:
					s.L.RaiseError("expected 2 or 3 arguments")
					return 0
				}
				return s.Push("ColorSequenceKeypoint", v)
			},
		},
	}
}
