package model

type Permission struct {
	PermissionID int64        `json:"permission_id" gorm:"primaryKey;column:permission_id"`
	Name         string       `json:"name" gorm:"column:name"`
	EnName       string       `json:"en_name" gorm:"column:en_name"`
	URL          string       `json:"url" gorm:"column:url"`
	Pid          int64        `json:"pid" gorm:"column:pid"`
	ResourceType string       `json:"resource_type" gorm:"column:resource_type"`
	ActionType   string       `json:"action_type" gorm:"column:action_type"`
	Icon         string       `json:"icon" gorm:"column:icon"`
	OpenInNewTab int64        `json:"open_in_new_tab" gorm:"column:open_in_new_tab"`
	Sort         int64        `json:"sort" gorm:"column:sort"`
	Description  string       `json:"description" gorm:"column:description"`
	IsBackstage  int64        `json:"is_backstage" gorm:"column:is_backstage"`
	Children     []Permission `json:"children" gorm:"-"` // 子节点
}

type PermissionCreateRequest struct {
	Name         string `json:"name" binding:"required"`
	EnName       string `json:"en_name" binding:"required"`
	URL          string `json:"url"`
	Pid          int64  `json:"pid"`
	ResourceType string `json:"resource_type"`
	ActionType   string `json:"action_type"`
	Icon         string `json:"icon"`
	OpenInNewTab int64  `json:"open_in_new_tab"`
	Sort         int64  `json:"sort"`
	Description  string `json:"description"`
	IsBackstage  int64  `json:"is_backstage"`
}

type PermissionUpdateRequest struct {
	PermissionID int64  `json:"permission_id" gorm:"primaryKey;column:permission_id"`
	Name         string `json:"name"`
	EnName       string `json:"en_name"`
	URL          string `json:"url"`
	Pid          int64  `json:"pid"`
	ResourceType string `json:"resource_type"`
	ActionType   string `json:"action_type"`
	Icon         string `json:"icon"`
	OpenInNewTab int64  `json:"open_in_new_tab"`
	Sort         int64  `json:"sort"`
	Description  string `json:"description"`
	IsBackstage  int64  `json:"is_backstage"`
}

func (Permission) TableName() string {
	return "permissions"
}
func (PermissionCreateRequest) TableName() string {
	return "permissions"
}
func (PermissionUpdateRequest) TableName() string {
	return "permissions"
}
