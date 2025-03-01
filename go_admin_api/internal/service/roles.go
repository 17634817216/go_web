package service

import (
	"context"
	"errors"
	"fmt"
	"go_admin_api/global"
	"go_admin_api/internal/model"
)

type RolesService struct{}

func (s *RolesService) UpdateRoles(ctx context.Context, Role *model.Roles, PermissionIds []int64, RoleId int64) error {
	err := global.App.DB.Model([]model.Roles{}).Where("role_id=?", RoleId).Updates(Role).Error
	if err != nil {
		return fmt.Errorf("更新角色数据失败：%v", err)
	} else {
		err := global.App.DB.Delete([]model.RolePermissions{}, RoleId).Error
		if err != nil {
			return fmt.Errorf("更新删除角色数据失败：%v", err)
		}
		UpdateData := make([]model.RolePermissions, 0, len(PermissionIds))
		for _, PermissionID := range PermissionIds {
			UpdateData = append(UpdateData, model.RolePermissions{
				RoleId:       RoleId,
				PermissionID: PermissionID,
			})
		}
		if err := global.App.DB.Create(&UpdateData).Error; err != nil {
			return fmt.Errorf("更新存储角色菜单失败：%v", err)
		}

	}
	return nil

}

func (s *RolesService) GetRoles(ctx context.Context) ([]model.Roles, error) {
	var roles []model.Roles
	if err := global.App.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *RolesService) DeleteRole(ctx context.Context, id int64) error {
	return global.App.DB.Delete(&model.Roles{}, id).Error
}

func (s *RolesService) CreateRolesPermissions(ctx context.Context, Role *model.Roles, permissionIDs []int64) error {
	var count int64
	if err := global.App.DB.Model(&model.Roles{}).Where("role_name = ?", Role.RoleName).Count(&count).Error; err != nil {
		return fmt.Errorf("检查角色名称失败: %v", err)
	}
	if count > 0 {
		return errors.New("该角色名称已存在，请查证后再试")
	}
	// 插入角色，获取自增ID
	if err := global.App.DB.Create(Role).Error; err != nil {
		return err
	} else {
		rolePermissions := make([]model.RolePermissions, 0, len(permissionIDs))
		for _, permissionID := range permissionIDs {
			rolePermissions = append(rolePermissions, model.RolePermissions{
				RoleId:       Role.RoleId,  // 使用刚创建的角色ID
				PermissionID: permissionID, // 对应的权限ID
			})
		}

		if err := global.App.DB.Create(&rolePermissions).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *RolesService) GetRolesPermissions(ctx context.Context, id int64) ([]model.Permission, error) {
	//var roles []model.Permission
	var rolespermis []model.RolePermissions
	var permis []model.Permission
	if err := global.App.DB.Model([]model.RolePermissions{}).Where("role_id=?", id).Find(&rolespermis).Error; err != nil {
		return nil, err
	}
	var PernisIds = make([]int64, len(rolespermis))
	for _, RoleData := range rolespermis {
		PernisIds = append(PernisIds, RoleData.PermissionID)
	}
	if err := global.App.DB.Model([]model.Permission{}).Where("permission_id in ?", PernisIds).Find(&permis).Error; err != nil {
		return nil, err
	}

	return s.buildRolePermissionTree(permis, 0), nil
}

func (s *RolesService) buildRolePermissionTree(permissions []model.Permission, pid int64) []model.Permission {
	var tree []model.Permission
	for _, item := range permissions {
		if item.Pid == pid {
			item.Children = s.buildRolePermissionTree(permissions, item.PermissionID)
			tree = append(tree, item)
		}
	}
	return tree
}
