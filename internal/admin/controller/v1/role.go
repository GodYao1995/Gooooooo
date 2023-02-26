package v1

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/cd-home/Goooooo/internal/admin/types"
	"github.com/cd-home/Goooooo/internal/admin/version"
	"github.com/cd-home/Goooooo/internal/domain"
	"github.com/cd-home/Goooooo/internal/pkg/errno"
	"github.com/cd-home/Goooooo/internal/pkg/middleware/auth"
	"github.com/cd-home/Goooooo/internal/pkg/middleware/permission"
	"github.com/cd-home/Goooooo/internal/pkg/middleware/tracer"
	"github.com/cd-home/Goooooo/internal/pkg/res"
	"github.com/cd-home/Goooooo/internal/pkg/session"
	"github.com/cd-home/Goooooo/pkg/xhttp"
	"github.com/cd-home/Goooooo/pkg/xtracer"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

type RoleController struct {
	logic domain.RoleLogicFace
	log   *zap.Logger
}

func NewRoleController(apiV1 *version.APIV1, log *zap.Logger, logic domain.RoleLogicFace, store *session.RedisStore,
	perm *casbin.Enforcer, xtracer *xtracer.XTracer) {
	ctl := &RoleController{
		logic: logic,
		log:   log.WithOptions(zap.Fields(zap.String("module", "RoleController"))),
	}
	// API version
	v1 := apiV1.Group.Group("/role").Use(tracer.Tracing(xtracer))

	// Need Authorization
	needAuth := v1.Use(auth.AuthMiddleware(store))
	{
		needAuth.GET("/list", ctl.ListRole)
		needAuth.POST("/move", ctl.MoveRole)
	}

	// Need Permission
	needPerm := needAuth.Use(permission.PermissionMiddleware(perm))
	{
		needPerm.POST("/create", ctl.CreateRole)
		needPerm.DELETE("/delete", ctl.DeleteRole)
		needPerm.PUT("/update", ctl.UpdateRole)
	}
}

// CreateRole
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept  json
// @Produce json
// @Param CreateRole body types.CreateRoleParam true "CreateRole"
// @Success 0 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Error"}
// @Router /create [POST]
func (r RoleController) CreateRole(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "RoleController-CreateRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("Controller", "CreateRole")
		span.Finish()
	}()
	params := types.CreateRoleParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := xhttp.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	err := r.logic.CreateRole(next, params.RoleName, params.RoleLevel, params.RoleIndex, params.Father)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.RoleCreateSuccess
	}
	resp.Success(ctx)
}

// DeleteRole
// @Summary Delete Role
// @Description Delete Role
// @Tags Role
// @Accept  json
// @Produce json
// @Param DeleteRole body types.DeleteRoleParam true "DeleteRole"
// @Success 0 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Error"}
// @Router /delete [DELETE]
func (r RoleController) DeleteRole(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "RoleController-DeleteRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("Controller", "DeleteRole")
		span.Finish()
	}()
	params := types.DeleteRoleParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := xhttp.ShouldBindQuery(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	err := r.logic.DeleteRole(next, params.RoleId)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Message = errno.Success
		resp.Code = 0
	}
	resp.Success(ctx)
}

// UpdateRole
// @Summary Update Role Name
// @Description Update Role Name
// @Tags Role
// @Accept  json
// @Produce json
// @Param UpdateRole body types.RenameRoleParam true "UpdateRole"
// @Success 0 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Error"}
// @Router /update [PUT]
func (r RoleController) UpdateRole(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "RoleController-UpdateRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("Controller", "UpdateRole")
		span.Finish()
	}()
	params := types.RenameRoleParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := xhttp.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	if err := r.logic.UpdateRole(next, params.RoleId, params.RoleName); err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.Success
	}
	resp.Success(ctx)
}

// ListRole
// @Summary List Role
// @Description List Role
// @Tags Role
// @Accept  json
// @Produce json
// @Param ListRole body types.ListRoleParam true "ListRole"
// @Success 0 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Error"}
// @Router /list [GET]
func (r RoleController) ListRole(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "RoleController-ListRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("Controller", "ListRole")
		span.Finish()
	}()
	params := types.ListRoleParam{}
	resp := res.CommonResponse{Code: 1}
	if ok, valid := xhttp.ShouldBind(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	if params.RoleLevel >= 2 && params.Father == nil {
		resp.Message = errno.ErrorNotEnoughParam.Error()
		resp.Failure(ctx)
		return
	}
	views, err := r.logic.RetrieveRoles(next, params.RoleLevel, params.Father)
	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.Success
		resp.Data = views
	}
	resp.Success(ctx)
}

// MoveRole
// @Summary Move Role
// @Description Move Role
// @Tags Role
// @Accept  json
// @Produce json
// @Param ListRole body types.ListRoleParam true "ListRole"
// @Success 0 {object} types.CommonResponse {"code":1,"data":null,"msg":"Success"}
// @Failure 1 {object} types.CommonResponse {"code":1,"data":null,"msg":"Error"}
// @Router /move [POST]
func (r RoleController) MoveRole(ctx *gin.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx.Request.Context(), "RoleController-MoveRole")
	next := opentracing.ContextWithSpan(context.Background(), span)
	defer func() {
		span.SetTag("Controller", "MoveRole")
		span.Finish()
	}()
	resp := res.CommonResponse{Code: 1}
	params := types.MoveRoleParam{}
	if ok, valid := xhttp.ShouldBindJSON(ctx, &params); !ok {
		resp.Message = valid
		resp.Failure(ctx)
		return
	}
	if err := r.logic.MoveRole(next, params.RoleId, params.Father); err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = 0
		resp.Message = errno.Success
	}
	resp.Success(ctx)
}
