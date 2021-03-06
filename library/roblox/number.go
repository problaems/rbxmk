package reflect

import (
	. "github.com/anaminus/rbxmk"
	"github.com/robloxapi/types"
	lua "github.com/yuin/gopher-lua"
)

func init() { register(Number) }
func Number() Reflector {
	return Reflector{
		Name: "number",
		PushTo: func(s State, r Reflector, v types.Value) (lvs []lua.LValue, err error) {
			return []lua.LValue{lua.LNumber(v.(types.Double))}, nil
		},
		PullFrom: func(s State, r Reflector, lvs ...lua.LValue) (v types.Value, err error) {
			if n, ok := lvs[0].(lua.LNumber); ok {
				return types.Double(n), nil
			}
			return nil, TypeError(nil, 0, "number")
		},
	}
}
