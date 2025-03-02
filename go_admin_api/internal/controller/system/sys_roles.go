package system

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go_admin_api/internal/model"
	"go_admin_api/internal/service"
	"go_admin_api/utils"
	"io"
	"strconv"
)

type RolesServiceController struct {
	roleService *service.RolesService
}

func NewRoleController() *RolesServiceController {
	return &RolesServiceController{
		roleService: &service.RolesService{},
	}
}

func (c *RolesServiceController) Create(ctx *gin.Context) {
	//body, _ := io.ReadAll(ctx.Request.Body)
	//ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // 重新设置请求体
	var request model.CreateRoleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.AdminFailed(ctx, "读取请求体失败")
		return
	}

	if err := c.roleService.CreateRolesPermissions(ctx, &request.Roles, request.PermissionIds); err != nil {
		utils.AdminFailed(ctx, err.Error())
		return
	}
	utils.OpensResponse(ctx, "添加菜单成功")

}

func (c *RolesServiceController) GetRoles(ctx *gin.Context) {
	data, err := c.roleService.GetRoles(ctx)
	if err != nil {
		utils.AdminFailed(ctx, err.Error())
		return
	}
	utils.SuccessData(ctx, data)
}

func (c *RolesServiceController) UpadteRole(ctx *gin.Context) {
	idStr := ctx.Param("role_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}
	body, _ := io.ReadAll(ctx.Request.Body)

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	var request model.CreateRoleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.AdminFailed(ctx, "请求体鉴权失败，请重试")
	}

	if err := c.roleService.UpdateRoles(ctx, &request.Roles, request.PermissionIds, id); err != nil {
		utils.AdminFailed(ctx, "更新数据失败，请稍后重试")
		return
	}
	utils.OpensResponse(ctx, "更新数据成功，请稍后查看")

}

func (c *RolesServiceController) DeleteRoles(ctx *gin.Context) {
	role_id := ctx.Param("role_id")
	id, err := strconv.ParseInt(role_id, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}
	if err := c.roleService.DeleteRole(ctx, id); err != nil {
		utils.AdminFailed(ctx, err.Error())
		return
	}
	utils.AdminFailed(ctx, "删除角色成功，请稍后查看")
}

func (c *RolesServiceController) GetRolesPermissions(ctx *gin.Context) {
	role_id := ctx.Param("role_id")
	id, err := strconv.ParseInt(role_id, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}
	rolespermis, err := c.roleService.GetRolesPermissions(ctx, id)
	if err != nil {
		utils.AdminFailed(ctx, err.Error())
		return
	}
	utils.SuccessData(ctx, rolespermis)
}
