package service

import (
	"context"
	"errors"
	"go_admin_api/global"
	"go_admin_api/internal/model"
)

type PermissionService struct{}

// CreatePermission 创建权限
func (s *PermissionService) CreatePermission(ctx context.Context, permission *model.PermissionCreateRequest) error {

	return global.App.DB.Create(permission).Error
}

// UpdatePermission 修改权限
func (s *PermissionService) UpdatePermission(ctx context.Context, permission *model.PermissionUpdateRequest, id int64) error {
	return global.App.DB.Model(permission).Updates(permission).Where("permission_id = ?", id).Error
}

// DeletePermission 删除子节点
func (s *PermissionService) DeletePermission(ctx context.Context, id int64) error {
	// 检查是否有子节点
	var count int64
	if err := global.App.DB.Model(&model.Permission{}).Where("pid = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该权限存在子节点，无法删除")
	}

	return global.App.DB.Delete(&model.Permission{}, id).Error
}

// GetPermissionTree 获取权限树
func (s *PermissionService) GetPermissionTree(ctx context.Context, is_backstage int64) ([]model.Permission, error) {
	var permissions []model.Permission

	// 获取所有权限
	if err := global.App.DB.Where("is_backstage = ?", is_backstage).Order("sort").Find(&permissions).Error; err != nil {
		return nil, err
	}

	// 构建树状结构
	return s.buildPermissionTree(permissions, 0), nil
}

func (s *PermissionService) buildPermissionTree(permissions []model.Permission, pid int64) []model.Permission {
	var tree []model.Permission

	for _, item := range permissions {
		if item.Pid == pid {
			item.Children = s.buildPermissionTree(permissions, item.PermissionID)
			tree = append(tree, item)
		}
	}

	return tree
}
