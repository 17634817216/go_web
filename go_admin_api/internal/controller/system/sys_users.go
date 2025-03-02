package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin_api/internal/model"
	"go_admin_api/internal/service"
	"go_admin_api/utils"
	"strconv"
)

type UserServiceController struct {
	userService *service.UsersService
}

func NewUserController() *UserServiceController {
	return &UserServiceController{
		userService: &service.UsersService{},
	}
}

func (s *UserServiceController) GetlistPage(ctx *gin.Context) {
	var query model.UserQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		//ctx.JSON(400, gin.H{"error": "参数错误"})
		utils.AdminFailed(ctx, "参数错误,请查证后再试")

		return
	}

	result, err := s.userService.GetUsersPages(ctx, query) // 修正字段名和方法名
	if err != nil {
		ctx.JSON(500, gin.H{ // 服务器错误状态码
			"error": err.Error(), // 错误信息转换
		})
		utils.AdminFailed(ctx, fmt.Sprintf("获取数据失败: %+v", err))
		return
	}
	utils.SuccessData(ctx, result)
	return

}

func (s *UserServiceController) Create(ctx *gin.Context) {
	var query model.User
	if err := ctx.ShouldBindJSON(&query); err != nil {
		//ctx.JSON(400, gin.H{"error": "参数错误"})
		utils.AdminFailed(ctx, "参数错误,请查证后再试")

		return
	}
	err := s.userService.CreateUsers(ctx, &query) // 修正字段名和方法名
	if err != nil {
		utils.AdminFailed(ctx, fmt.Sprintf("添加用户信息失败: %+v", err))
		return
	}
	utils.OpensResponse(ctx, "添加用户信息成功")
	return

}

func (s *UserServiceController) Update(ctx *gin.Context) {
	idStr := ctx.Param("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}
	var query model.UpdateUser
	if err := ctx.ShouldBindJSON(&query); err != nil {
		//ctx.JSON(400, gin.H{"error": "参数错误"})
		utils.AdminFailed(ctx, "参数错误,请查证后再试")

		return
	}
	errs := s.userService.UpdateUsers(ctx, &query, id) // 修正字段名和方法名
	if errs != nil {
		utils.AdminFailed(ctx, fmt.Sprintf("修改用户信息失败: %+v", errs))
		return
	}
	utils.OpensResponse(ctx, "修改用户信息成功")
	return

}

func (s *UserServiceController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "无效的 ID"})
		return
	}
	errs := s.userService.DeleteUsers(ctx, id) // 修正字段名和方法名
	if errs != nil {
		utils.AdminFailed(ctx, fmt.Sprintf("删除用户信息失败: %+v", errs))
		return
	}
	utils.OpensResponse(ctx, "删除用户信息成功")
	return

}
