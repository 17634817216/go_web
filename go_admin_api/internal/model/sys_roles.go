package model

type Roles struct {
	RoleId      int64  `json:"role_id" gorm:"primaryKey;column:role_id"`
	RoleName    string `json:"role_name" gorm:"column:role_name" v:"required#角色名称不能为空"`
	IsBackstage int64  `json:"is_backstage" gorm:"column:is_backstage"  v:"required#前后台权限不能为空"`
	Description string `json:"description" gorm:"column:description"`
}
type CreateRoleRequest struct {
	Roles         Roles   `json:"roles"`
	PermissionIds []int64 `json:"permission_ids"`
}

type RolePermissions struct {
	RoleId       int64 `json:"role_id" gorm:"primaryKey;column:role_id"`
	PermissionID int64 `json:"permission_id" gorm:"column:permission_id"` // 权限ID数组

}

func (Roles) TableName() string {
	return "roles"
}
func (RolePermissions) TableName() string {
	return "role_permissions"
}
