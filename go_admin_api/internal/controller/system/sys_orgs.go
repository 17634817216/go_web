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

type OrganizationController struct {
	roleService *service.OrganizationService
}

func NewOrganizationController() *OrganizationController {
	return &OrganizationController{
		roleService: &service.OrganizationService{},
	}
}

func (c *OrganizationController) Create(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)
	log.Printf("Received body: %s", string(body))
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body)) // 重新设置请求体
	var organization model.CreateOrganization
	if err := ctx.ShouldBindJSON(&organization); err != nil {
		utils.AdminFailed(ctx, err.Error())

		return
	}

	fmt.Println(organization)
	if err := c.roleService.CreateOrganization(ctx, &organization); err != nil {
		utils.AdminFailed(ctx, err.Error())

		return
	}
	utils.OpensResponse(ctx, "添加组织机构成功")

}

func (c *OrganizationController) Update(ctx *gin.Context) {
	orgIdStr := ctx.Param("org_id")
	orgId, err := strconv.ParseInt(orgIdStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}
	var organization model.UpdateOrganization
	if err := ctx.ShouldBindJSON(&organization); err != nil {
		utils.AdminFailed(ctx, err.Error())

		return
	}
	fmt.Println(organization)
	if err := c.roleService.UpdateOrganization(ctx, &organization, orgId); err != nil {
		utils.AdminFailed(ctx, err.Error())

		return
	}
	utils.OpensResponse(ctx, "修改组织机构成功")

}

func (c *OrganizationController) GetOrgTree(ctx *gin.Context) {
	tree, err := c.roleService.GetOrganization(ctx)
	if err != nil {
		utils.AdminFailed(ctx, err.Error())

		return
	}
	utils.SuccessData(ctx, tree)

}

func (c *OrganizationController) DeleteOrg(ctx *gin.Context) {
	orgIdStr := ctx.Param("org_id")
	orgId, err := strconv.ParseInt(orgIdStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}
	if err := c.roleService.DeleteOrganization(ctx, orgId); err != nil {
		utils.AdminFailed(ctx, err.Error())

		return
	}
	utils.OpensResponse(ctx, "删除组织信息成功")

}
