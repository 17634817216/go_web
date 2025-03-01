package system

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin_api/internal/model"
	"go_admin_api/internal/service"
	"go_admin_api/utils"
	"io"
	"log"
	"strconv"
)

type PermissionController struct {
	permissionService *service.PermissionService
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		permissionService: &service.PermissionService{},
	}
}

func (c *PermissionController) Create(ctx *gin.Context) {
	//body, _ := io.ReadAll(ctx.Request.Body)
	//fmt.Println(body)
	body, _ := io.ReadAll(ctx.Request.Body)
	log.Printf("Received body: %s", string(body))
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // 重新设置请求体
	var permission model.PermissionCreateRequest
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		utils.AdminFailed(ctx, err.Error())

		return
	}

	fmt.Println(permission)
	if err := c.permissionService.CreatePermission(ctx, &permission); err != nil {
		utils.AdminFailed(ctx, err.Error())

		return
	}
	utils.OpensResponse(ctx, "添加菜单成功")

	//ctx.JSON(200, gin.H{"data": permission})
}

// Update 更新权限
func (c *PermissionController) Update(ctx *gin.Context) {
	idStr := ctx.Param("permission_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}

	var permission model.PermissionUpdateRequest

	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.permissionService.UpdatePermission(ctx, &permission, id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "更新成功"})
}

// Delete 删除权限
func (c *PermissionController) Delete(ctx *gin.Context) {
	// 获取路径参数 id
	idStr := ctx.Param("id")

	// 将 id 从 string 转换为 int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}

	// 调用服务层方法删除权限
	if err := c.permissionService.DeletePermission(ctx, id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	ctx.JSON(200, gin.H{"message": "删除成功"})
}

// GetTree 获取权限树
func (c *PermissionController) GetTree(ctx *gin.Context) {
	isbackstageStr := ctx.Param("is_backstage")

	// 将 id 从 string 转换为 int64
	isbackstage, err := strconv.ParseInt(isbackstageStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 状态获取"})
		return
	}
	tree, err := c.permissionService.GetPermissionTree(ctx, isbackstage)
	if err != nil {
		//ctx.JSON(500, gin.H{"error": err.Error()})
		utils.AdminFailed(ctx, err.Error())

		return
	}
	utils.SuccessData(ctx, tree)
	//ctx.JSON(200, gin.H{"data": tree})
}
