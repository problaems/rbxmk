package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(CallbackDesc) }
func CallbackDesc() Reflector {
	return Reflector{
		Name:     "CallbackDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Metatable: Metatable{
			"__eq": func(s State) int {
				v := s.Pull(1, "CallbackDesc").(rtypes.CallbackDesc)
				op := s.Pull(2, "CallbackDesc").(rtypes.CallbackDesc)
				s.L.Push(lua.LBool(v == op))
				return 1
			},
		},
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.CallbackDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
			},
			"Parameters": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.CallbackDesc)
				array := make(rtypes.Array, len(desc.Parameters))
				for i, param := range desc.Parameters {
					p := param
					array[i] = rtypes.ParameterDesc{Parameter: p}
				}
				return s.Push(array)
			}},
			"SetParameters": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.CallbackDesc)
				array := s.Pull(2, "Array").(rtypes.Array)
				params := make([]rbxdump.Parameter, len(array))
				for i, paramDesc := range array {
					param, ok := paramDesc.(rtypes.ParameterDesc)
					if !ok {
						TypeError(s.L, 3, param.Type())
					}
					params[i] = param.Parameter
				}
				desc.Parameters = params
				return 0
			}},
			"ReturnType": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					returnType := desc.ReturnType
					return s.Push(rtypes.TypeDesc{Embedded: returnType})
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.CallbackDesc)
					desc.ReturnType = s.Pull(3, "TypeDesc").(rtypes.TypeDesc).Embedded
				},
			},
			"Security": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.CallbackDesc)
					return s.Push(types.String(desc.Security))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.CallbackDesc)
					desc.Security = string(s.Pull(3, "string").(types.String))
				},
			},
			"Tag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.CallbackDesc)
				tag := string(s.Pull(2, "string").(types.String))
				return s.Push(types.Bool(desc.GetTag(tag)))
			}},
			"Tags": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.CallbackDesc)
				tags := desc.GetTags()
				array := make(rtypes.Array, len(tags))
				for i, tag := range tags {
					array[i] = types.String(tag)
				}
				return s.Push(array)
			}},
			"SetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.CallbackDesc)
				tags := make([]string, s.Count()-1)
				for i := 2; i <= s.Count(); i++ {
					tags[i-2] = string(s.Pull(i, "string").(types.String))
				}
				desc.SetTag(tags...)
				return 0
			}},
			"UnsetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.CallbackDesc)
				tags := make([]string, s.Count()-1)
				for i := 2; i <= s.Count(); i++ {
					tags[i-2] = string(s.Pull(i, "string").(types.String))
				}
				desc.UnsetTag(tags...)
				return 0
			}},
		},
	}
}
