package service

import (
	"context"
	"fmt"
	"go_admin_api/global"
	"go_admin_api/internal/model"
	"go_admin_api/utils"
	"math"
)

type UsersService struct{}

func (s *UsersService) CreateUsers(ctx context.Context, users *model.User) error {
	var usercount int64
	global.App.DB.Model(users).Select("user_id").Where("username=?", users.Username).Count(&usercount)
	if usercount > 0 {
		return fmt.Errorf("用户名 %s 已存在", users.Username)
	}
	global.App.DB.Model(users).Where("nickname=?", users.Nickname).Count(&usercount)
	if usercount > 0 {
		return fmt.Errorf("昵称 %s 已存在", users.Nickname)
	}
	return global.App.DB.Create(users).Error
}

func (s *UsersService) UpdateUsers(ctx context.Context, users *model.UpdateUser, ID int64) error {
	fmt.Println(ID)
	err := global.App.DB.Model([]model.UpdateUser{}).Where("user_id=?", ID).Updates(users).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *UsersService) GetUsersPages(ctx context.Context, query model.UserQuery) (*model.Pagination, error) {
	//var usersdata model.Pagination
	db := global.App.DB.Model([]model.GetUser{})
	if query.Search != "" {
		search := "%" + query.Search + "%"
		db = db.Where(
			"username LIKE ? OR "+
				"nickname LIKE ? OR "+
				"leader_name LIKE ? OR "+
				"email LIKE ? OR "+
				"mobile_phone LIKE ?",
			search, search, search, search, search,
		)
	}

	if query.OrgID != 0 {
		db = db.Where(
			"org_id = ? ",
			query.OrgID,
		)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, fmt.Errorf("查询总数失败: %v", err)
	}

	var users []model.GetUser
	err := db.Scopes(utils.Paginate(query.Page, query.PageSize)).
		Order("update_time DESC").
		Find(&users).Error

	if err != nil {
		return nil, fmt.Errorf("数据查询失败: %v", err)
	}

	userdata := s.GetOrgData(users)

	//global.App.DB.Model()
	fmt.Println(float64(total))
	fmt.Println((int(math.Ceil(float64(total) / float64(query.PageSize)))) + 1)
	response := &model.Pagination{
		Page:       query.Page,
		PageSize:   query.PageSize,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(query.PageSize))),
		Data:       userdata,
	}

	return response, nil
}

func (s *UsersService) DeleteUsers(ctx context.Context, ID int64) error {
	return global.App.DB.Model([]model.User{}).Delete("user_id=?", ID).Error
}

func (s *UsersService) GetOrgData(userpermissions []model.GetUser) []model.GetOrgUser {
	var roleSystem = make(map[int64]model.Orglist)
	var orgData []model.Orglist
	var orgsData = make([]model.GetOrgUser, len(orgData))

	result := global.App.DB.Model(&[]model.Orglist{}).Find(&orgData)
	if result.Error != nil {
		return nil
	}
	for _, item := range orgData {
		roleSystem[item.OrgId] = model.Orglist{
			OrgId:       item.OrgId,
			Name:        item.Name,
			ParentId:    item.ParentId,
			Level:       item.Level,
			Description: item.Description,
		}
	}
	for _, item := range userpermissions {
		var hierarchy model.OrgHierarchy
		currentID := item.OrgID
		maxLevel := 4 // 最大追溯层级

		for i := 0; i < maxLevel; i++ {
			orgInfo, exists := roleSystem[currentID]
			if !exists {
				//log.Printf("组织ID %d 不存在，用户ID: %d", currentID, item.UserID)
				break
			}

			// 根据层级填充信息
			switch orgInfo.Level {
			case 4:
				hierarchy.LineName = orgInfo.Name
				hierarchy.LineID = orgInfo.OrgId
			case 3:
				hierarchy.DepartmentName = orgInfo.Name
				hierarchy.DepartmentID = orgInfo.OrgId
			case 2:
				hierarchy.SecondaryFactoryName = orgInfo.Name
				hierarchy.SecondaryFactoryID = orgInfo.OrgId
			case 1:
				hierarchy.FactoryName = orgInfo.Name
				hierarchy.FactoryID = orgInfo.OrgId
			}

			// 向上追溯父级组织
			if orgInfo.ParentId == 0 {
				break
			}
			currentID = orgInfo.ParentId
		}
		orgsData = append(orgsData, model.GetOrgUser{
			UserID:                item.UserID,
			Username:              item.Username,
			Password:              item.Password,
			MobilePhone:           item.MobilePhone,
			Email:                 item.Email,
			Status:                item.Status,
			OrgID:                 item.OrgID,
			Factoryname:           hierarchy.FactoryName,
			Factoryid:             hierarchy.FactoryID,
			SecondaryFactoryyname: hierarchy.SecondaryFactoryName,
			SecondaryFactoryy:     hierarchy.SecondaryFactoryID,
			Linename:              hierarchy.LineName,
			Lineid:                hierarchy.LineID,
			Departmentname:        hierarchy.DepartmentName,
			Departmentid:          hierarchy.DepartmentID,
			RoleIDs:               item.RoleIDs,
			CreateTime:            item.CreateTime,
			UpdateTime:            item.UpdateTime,
			WechatUserID:          item.WechatUserID,
			WechatChatID:          item.WechatChatID,
			EntryNumber:           item.EntryNumber,
			NotificationSettings:  item.NotificationSettings,
			LoggingStatus:         item.LoggingStatus,
			Nickname:              item.Nickname,
			LeaderName:            item.LeaderName,
		})
	}

	return orgsData
}
