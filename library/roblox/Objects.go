package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Objects) }
func Objects() Reflector {
	return Reflector{
		Name: "Objects",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			objects, ok := v.(rtypes.Objects)
			if !ok {
				return nil, TypeError(nil, 0, "Objects")
			}
			instRfl := s.Reflector("Instance")
			table := s.L.CreateTable(len(objects), 0)
			for i, v := range objects {
				lv, err := instRfl.PushTo(s, instRfl, v)
				if err != nil {
					return nil, err
				}
				table.RawSetInt(i+1, lv[0])
			}
			return []lua.LValue{table}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			table, ok := lvs[0].(*lua.LTable)
			if !ok {
				return nil, TypeError(nil, 0, "table")
			}
			instRfl := s.Reflector("Instance")
			n := table.Len()
			objects := make(rtypes.Objects, n)
			for i := 1; i <= n; i++ {
				v, err := instRfl.PullFrom(s, instRfl, table.RawGetInt(i))
				if err != nil {
					return nil, err
				}
				objects[i-1] = v.(*rtypes.Instance)
			}
			return objects, nil
		},
	}
}
