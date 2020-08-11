package rtypes

import (
	"github.com/robloxapi/types"
)

type Array []types.Value

func (Array) Type() string { return "Array" }

type Dictionary map[string]types.Value

func (Dictionary) Type() string { return "Dictionary" }

type Tuple []types.Value

func (Tuple) Type() string { return "Tuple" }

type Objects []*Instance

func (Objects) Type() string { return "Objects" }
