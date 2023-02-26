package logic

import (
	v1 "github.com/cd-home/Goooooo/internal/admin/logic/v1"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	v1.NewUserLogic,
	v1.NewDirectoryrLogic,
	v1.NewFileLogic,
	v1.NewRoleLogic,
)
