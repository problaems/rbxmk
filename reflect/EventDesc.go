package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/types"
)

func EventDesc() Type {
	return Type{
		Name:     "EventDesc",
		PushTo:   PushTypeTo,
		PullFrom: PullTypeFrom,
		Members: Members{
			"Name": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
					return s.Push(types.String(desc.Name))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EventDesc)
					desc.Name = string(s.Pull(3, "string").(types.String))
				},
			},
			"Parameters": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EventDesc)
				array := make(rtypes.Array, len(desc.Parameters))
				for i, param := range desc.Parameters {
					array[i] = rtypes.ParameterDesc{Parameter: param}
				}
				return s.Push(array)
			}},
			"SetParameters": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EventDesc)
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
			"Security": Member{
				Get: func(s State, v types.Value) int {
					desc := v.(rtypes.EventDesc)
					return s.Push(types.String(desc.Security))
				},
				Set: func(s State, v types.Value) {
					desc := v.(rtypes.EventDesc)
					desc.Security = string(s.Pull(3, "string").(types.String))
				},
			},
			"Tag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EventDesc)
				tag := string(s.Pull(2, "string").(types.String))
				return s.Push(types.Bool(desc.GetTag(tag)))
			}},
			"Tags": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EventDesc)
				tags := desc.GetTags()
				array := make(rtypes.Array, len(tags))
				for i, tag := range tags {
					array[i] = types.String(tag)
				}
				return s.Push(array)
			}},
			"SetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EventDesc)
				tags := make([]string, s.Count()-1)
				for i := 2; i <= s.Count(); i++ {
					tags[i-2] = string(s.Pull(i, "string").(types.String))
				}
				desc.SetTag(tags...)
				return 0
			}},
			"UnsetTag": Member{Method: true, Get: func(s State, v types.Value) int {
				desc := v.(rtypes.EventDesc)
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