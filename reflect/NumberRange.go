package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	"github.com/yuin/gopher-lua"
)

func NumberRange() Type {
	return Type{
		Name:        "NumberRange",
		ReflectTo:   ReflectTypeTo,
		ReflectFrom: ReflectTypeFrom,
		Metatable: Metatable{
			"__tostring": func(s State, v types.Value) int {
				s.L.Push(lua.LString(v.(types.NumberRange).String()))
				return 1
			},
			"__eq": func(s State, v types.Value) int {
				op := s.Pull(2, "NumberRange").(types.NumberRange)
				return s.Push("bool", types.Bool(v.(types.NumberRange) == op))
			},
		},
		Members: map[string]Member{
			"Min": {Get: func(s State, v types.Value) int {
				return s.Push("float", types.Float(v.(types.NumberRange).Min))
			}},
			"Max": {Get: func(s State, v types.Value) int {
				return s.Push("float", types.Float(v.(types.NumberRange).Max))
			}},
		},
		Constructors: Constructors{
			"new": func(s State) int {
				var v types.NumberRange
				switch s.Count() {
				case 1:
					v.Min = float32(s.Pull(1, "float").(types.Float))
					v.Max = v.Min
				case 2:
					v.Min = float32(s.Pull(1, "float").(types.Float))
					v.Max = float32(s.Pull(2, "float").(types.Float))
				default:
					s.L.RaiseError("expected 1 or 2 arguments")
					return 0
				}
				return s.Push("NumberRange", v)
			},
		},
	}
}
