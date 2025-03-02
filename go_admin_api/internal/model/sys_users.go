package model

import (
	"gorm.io/datatypes"
	"time"
)

type User struct {
	UserID      int64          `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	Username    string         `gorm:"type:varchar(50);unique;not null" json:"username" validate:"required#用户名不能为空"`
	Password    string         `gorm:"type:varchar(255);not null" json:"password" validate:"required#请输入密码"`
	MobilePhone *string        `gorm:"type:varchar(20)" json:"mobile_phone,omitempty" validate:"omitempty,numeric,len=11"`
	Email       *string        `gorm:"type:varchar(100)" json:"email,omitempty" validate:"omitempty,email"`
	Status      int64          `gorm:"type:int;not null;default:1" json:"status"`
	OrgID       int64          `gorm:"type:int;not null" json:"org_id"`
	RoleIDs     datatypes.JSON `gorm:"type:json;not null" json:"role_ids"`
	CreateTime  time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"create_time"`

	WechatUserID         *string        `gorm:"type:varchar(200)" json:"wechat_user_id,omitempty"`
	WechatChatID         *string        `gorm:"type:varchar(200)" json:"wechat_chat_id,omitempty"`
	EntryNumber          *string        `gorm:"type:varchar(50)" json:"entry_number,omitempty"`
	NotificationSettings datatypes.JSON `gorm:"type:json" json:"notification_settings,omitempty"`
	LoggingStatus        *int64         `gorm:"type:int" json:"logging_status,omitempty"`
	Nickname             *string        `gorm:"type:varchar(100)" json:"nickname,omitempty"`
	LeaderName           *string        `gorm:"type:varchar(100)" json:"leader_name,omitempty"`
}

type UpdateUser struct {
	UserID               int64          `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	Username             string         `gorm:"type:varchar(50);unique;not null" json:"username" validate:"required#用户名不能为空"`
	Password             string         `gorm:"type:varchar(255);not null" json:"password" validate:"required#请输入密码"`
	MobilePhone          *string        `gorm:"type:varchar(20)" json:"mobile_phone,omitempty" validate:"omitempty,numeric,len=11"`
	Email                *string        `gorm:"type:varchar(100)" json:"email,omitempty" validate:"omitempty,email"`
	Status               int64          `gorm:"type:int;not null;default:1" json:"status"`
	OrgID                int64          `gorm:"type:int;not null" json:"org_id"`
	RoleIDs              datatypes.JSON `gorm:"type:json;not null" json:"role_ids"`
	UpdateTime           time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"update_time"`
	WechatUserID         *string        `gorm:"type:varchar(200)" json:"wechat_user_id,omitempty"`
	WechatChatID         *string        `gorm:"type:varchar(200)" json:"wechat_chat_id,omitempty"`
	EntryNumber          *string        `gorm:"type:varchar(50)" json:"entry_number,omitempty"`
	NotificationSettings datatypes.JSON `gorm:"type:json" json:"notification_settings,omitempty"`
	LoggingStatus        *int64         `gorm:"type:int" json:"logging_status,omitempty"`
	Nickname             *string        `gorm:"type:varchar(100)" json:"nickname,omitempty"`
	LeaderName           *string        `gorm:"type:varchar(100)" json:"leader_name,omitempty"`
}
type GetUser struct {
	UserID               int64          `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	Username             string         `gorm:"type:varchar(50);unique;not null" json:"username" `
	Password             string         `gorm:"type:varchar(255);not null" json:"password" `
	MobilePhone          *string        `gorm:"type:varchar(20)" json:"mobile_phone" validate:"omitempty,numeric,len=11"`
	Email                *string        `gorm:"type:varchar(100)" json:"email" validate:"omitempty,email"`
	Status               int64          `gorm:"type:int;not null;default:1" json:"status"`
	OrgID                int64          `gorm:"type:int;not null" json:"org_id"`
	RoleIDs              datatypes.JSON `gorm:"type:json;not null" json:"role_ids"`
	CreateTime           time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime           time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"update_time"`
	WechatUserID         *string        `gorm:"type:varchar(200)" json:"wechat_user_id"`
	WechatChatID         *string        `gorm:"type:varchar(200)" json:"wechat_chat_id"`
	EntryNumber          *string        `gorm:"type:varchar(50)" json:"entry_number"`
	NotificationSettings datatypes.JSON `gorm:"type:json" json:"notification_settings"`
	LoggingStatus        *int64         `gorm:"type:int" json:"logging_status"`
	Nickname             *string        `gorm:"type:varchar(100)" json:"nickname"`
	LeaderName           *string        `gorm:"type:varchar(100)" json:"leader_name"`
}
type UserQuery struct {
	Page     int    `form:"page" json:"page"`           // 页码
	PageSize int    `form:"page_size" json:"page_size"` // 每页数量
	Search   string `form:"search" json:"search"`       // 用户名模糊查询
	OrgID    int64  `form:"org_id" json:"org_id"`       // 组织ID精确查询

}

type Pagination struct {
	Page       int         `json:"page"`       // 当前页码
	PageSize   int         `json:"pageSize"`   // 每页数量
	Total      int64       `json:"total"`      // 总记录数
	TotalPages int         `json:"totalPages"` // 总页数
	Data       interface{} `json:"data"`       // 数据列表
}

type GetOrgUser struct {
	UserID                int64          `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	Username              string         `gorm:"type:varchar(50);unique;not null" json:"username" validate:"required,alphanumunicode"`
	Password              string         `gorm:"type:varchar(255);not null" json:"password" validate:"required"`
	MobilePhone           *string        `gorm:"type:varchar(20)" json:"mobile_phone,omitempty" validate:"omitempty,numeric,len=11"`
	Email                 *string        `gorm:"type:varchar(100)" json:"email,omitempty" validate:"omitempty,email"`
	Status                int64          `gorm:"type:int;not null;default:1" json:"status"`
	OrgID                 int64          `gorm:"type:int;not null" json:"org_id"`
	Factoryname           string         `json:"factory_name" validate:"required"`
	Factoryid             int64          `json:"factory_id" validate:"required"`
	SecondaryFactoryyname string         ` json:"secondary_factoryy_name" validate:"required"`
	SecondaryFactoryy     int64          `json:"secondary_factoryy_id" validate:"required"`
	Linename              string         `json:"line_name" validate:"required"`
	Lineid                int64          `json:"line_id" validate:"required"`
	Departmentname        string         ` json:"department_name" validate:"required"`
	Departmentid          int64          `json:"department_id" validate:"required"`
	RoleIDs               datatypes.JSON `gorm:"type:json;not null" json:"role_ids"`
	CreateTime            time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime            time.Time      `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"update_time"`
	WechatUserID          *string        `gorm:"type:varchar(200)" json:"wechat_user_id,omitempty"`
	WechatChatID          *string        `gorm:"type:varchar(200)" json:"wechat_chat_id,omitempty"`
	EntryNumber           *string        `gorm:"type:varchar(50)" json:"entry_number,omitempty"`
	NotificationSettings  datatypes.JSON `gorm:"type:json" json:"notification_settings,omitempty"`
	LoggingStatus         *int64         `gorm:"type:int" json:"logging_status,omitempty"`
	Nickname              *string        `gorm:"type:varchar(100)" json:"nickname,omitempty"`
	LeaderName            *string        `gorm:"type:varchar(100)" json:"leader_name,omitempty"`
}

type OrgHierarchy struct {
	LineName             string
	LineID               int64
	DepartmentName       string
	DepartmentID         int64
	SecondaryFactoryName string
	SecondaryFactoryID   int64
	FactoryName          string
	FactoryID            int64
}

type Orglist struct {
	OrgId       int64  `json:"org_id" gorm:"primaryKey;column:org_id"` // 使用 org_id 而不是 permission_id
	Name        string `json:"name" gorm:"column:name" v:"required#名称不能为空"`
	ParentId    int64  `json:"parent_id" gorm:"column:parent_id"`
	Level       int64  `json:"level" gorm:"column:level"`
	Description string `json:"description" gorm:"column:description"`
}

func (User) TableName() string {
	return "users"
}
func (UpdateUser) TableName() string {
	return "users"
}

func (GetUser) TableName() string {
	return "users"
}
func (Orglist) TableName() string {
	return "organization"
}
