package xes

import (
	v1 "github.com/GodYao1995/Goooooo/internal/admin/xes/v1"
	"go.uber.org/fx"
)

var Module = fx.Provide(v1.NewUserEs)
