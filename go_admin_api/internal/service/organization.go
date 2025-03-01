package service

import (
	"context"
	"go_admin_api/global"
	"go_admin_api/internal/model"
)

type OrganizationService struct{}

func (s *OrganizationService) CreateOrganization(ctx context.Context, organization *model.CreateOrganization) error {

	return global.App.DB.Create(organization).Error
}

func (s *OrganizationService) UpdateOrganization(ctx context.Context, organization *model.UpdateOrganization, id int64) error {
	return global.App.DB.Model(organization).Updates(organization).Where("org_id=?", id).Error
}

func (s *OrganizationService) GetOrganization(ctx context.Context) ([]model.Organization, error) {
	var orgData []model.Organization
	err := global.App.DB.Model([]model.Organization{}).Order("sort").Find(&orgData).Error
	if err != nil {
		return nil, err
	}

	return s.buildOrganizationTree(orgData, 0), nil
}

func (s *OrganizationService) buildOrganizationTree(permissions []model.Organization, pid int64) []model.Organization {
	var tree []model.Organization

	for _, item := range permissions {
		if item.ParentId == pid {
			item.Children = s.buildOrganizationTree(permissions, item.OrgId)
			tree = append(tree, item)
		}
	}

	return tree
}

func (s *OrganizationService) DeleteOrganization(cex context.Context, id int64) error {
	err := global.App.DB.Model([]model.Organization{}).Delete("org_id=?", id).Error
	if err != nil {
		return err
	}
	return nil
}
