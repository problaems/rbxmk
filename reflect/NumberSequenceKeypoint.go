package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func NumberSequenceKeypoint() Type {
	return Type{
		Name:        "NumberSequenceKeypoint",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v Value) int {
				s.L.Push(lua.LString(v.(types.NumberSequenceKeypoint).String()))
				return 1
			},
			"__eq": func(s State, v Value) int {
				op := s.Pull(2, "NumberSequenceKeypoint").(types.NumberSequenceKeypoint)
				return s.Push("bool", types.Bool(v.(types.NumberSequenceKeypoint) == op))
			},
		},
		Members: map[string]Member{
			"Time": {Get: func(s State, v Value) int {
				return s.Push("float", types.Float(v.(types.NumberSequenceKeypoint).Time))
			}},
			"Value": {Get: func(s State, v Value) int {
				return s.Push("float", types.Float(v.(types.NumberSequenceKeypoint).Value))
			}},
			"Envelope": {Get: func(s State, v Value) int {
				return s.Push("float", types.Float(v.(types.NumberSequenceKeypoint).Envelope))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.NumberSequenceKeypoint
				switch s.Count() {
				case 2:
					v.Time = float32(s.Pull(1, "float").(types.Float))
					v.Value = float32(s.Pull(2, "float").(types.Float))
				case 3:
					v.Time = float32(s.Pull(1, "float").(types.Float))
					v.Value = float32(s.Pull(2, "float").(types.Float))
					v.Envelope = float32(s.Pull(3, "float").(types.Float))
				default:
					s.L.RaiseError("expected 2 or 3 arguments")
					return 0
				}
				return s.Push("NumberSequenceKeypoint", v)
			},
		},
	}
}
