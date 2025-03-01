package model

import "time"

// Organization 结构体定义
type Organization struct {
	OrgId       int64          `json:"org_id" gorm:"primaryKey;column:org_id"` // 使用 org_id 而不是 permission_id
	Name        string         `json:"name" gorm:"column:name" v:"required#名称不能为空"`
	ParentId    int64          `json:"parent_id" gorm:"column:parent_id"`
	Level       int64          `json:"level" gorm:"column:level"`
	Description string         `json:"description" gorm:"column:description"`
	Sort        int64          `json:"sort" gorm:"column:sort" v:"required#排序不能为空"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at"` // 使用 time.Time 类型
	UpdatedAt   time.Time      `json:"updated_at" gorm:"column:updated_at"` // 使用 time.Time 类型
	Children    []Organization `json:"children" gorm:"-"`                   // 子节点，类型改为 Organization
}

type CreateOrganization struct {
	Name        string    `json:"name" gorm:"column:name" v:"required#名称不能为空"`
	ParentId    int64     `json:"parent_id" gorm:"column:parent_id"`
	Level       int64     `json:"level" gorm:"column:level"`
	Description string    `json:"description" gorm:"column:description"`
	Sort        int64     `json:"sort" gorm:"column:sort" v:"required#排序不能为空"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"` // 使用 time.Time 类型

}

type UpdateOrganization struct {
	OrgId       int64     `json:"org_id" gorm:"primaryKey;column:org_id"` // 使用 org_id 而不是 permission_id
	Name        string    `json:"name" gorm:"column:name" v:"required#名称不能为空"`
	ParentId    int64     `json:"parent_id" gorm:"column:parent_id"`
	Level       int64     `json:"level" gorm:"column:level"`
	Description string    `json:"description" gorm:"column:description"`
	Sort        int64     `json:"sort" gorm:"column:sort" v:"required#排序不能为空"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"` // 使用 time.Time 类型
}

func (Organization) TableName() string {
	return "organization"
}

func (UpdateOrganization) TableName() string {
	return "organization"
}

func (CreateOrganization) TableName() string {
	return "organization"
}
