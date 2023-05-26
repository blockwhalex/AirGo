package service

import (
	"AirGo/global"
	"AirGo/model"
	"log"
)

// 根据uId查角色Ids
func FindRoleIdsByuId(uId int) ([]int, error) {
	type Roles struct {
		RoleID int
	}
	var RolesSlice []Roles
	err := global.DB.Model(&model.UserAndRole{}).Where("user_id=?", uId).Find(&RolesSlice).Error
	if err != nil {
		log.Println("DB err:", err)
		return nil, err
	}
	//角色id 数组
	var roleIds []int
	for item := range RolesSlice {
		roleIds = append(roleIds, RolesSlice[item].RoleID)
	}
	return roleIds, nil
}

// 分页获取角色列表
func GetRoleList(roleParams *model.PaginationParams) ([]model.Role, error) {
	//log.Println("接收role params：", roleParams.Search)
	var roles []model.Role
	if roleParams.Search != "" {
		//gorm 模糊搜索
		err := global.DB.Debug().Where("role_name like ?", ("%" + roleParams.Search + "%")).Preload("Menus").Limit(roleParams.PageSize).Offset((roleParams.PageNum - 1) * roleParams.PageSize).Find(&roles).Error
		return roles, err
	}
	err := global.DB.Preload("Menus").Limit(roleParams.PageSize).Offset((roleParams.PageNum - 1) * roleParams.PageSize).Find(&roles).Error
	return roles, err
}

// 修改角色信息
func ModifyRoleInfo(roleInfo *model.Role) error {
	//先更新角色信息
	err := global.DB.Save(&roleInfo).Error
	if err != nil {
		return err
	}
	//再更新关联信息
	var routerSlice []model.DynamicRoute
	global.DB.Where("id in (?)", roleInfo.Nodes).Find(&routerSlice)
	err = global.DB.Model(&roleInfo).Association("Menus").Replace(&routerSlice)
	return err
}

// 新建角色
func AddRole(role *model.Role) error {
	//先查关联的动态路由
	//var routerSlice []model.SysDynamicRouter
	err := global.DB.Where("id in (?)", role.Nodes).Find(&role.Menus).Error
	if err != nil {
		return err
	}
	err = global.DB.Create(&role).Error
	return err
}

// 删除角色
func DelRole(id int) error {
	var role model.Role
	//查询全部关联并删除
	err := global.DB.Debug().Where("id = ?", id).Preload("Menus").Find(&role).Error
	if err != nil {
		return err
	}
	//fmt.Println("查到的role：", role)
	global.DB.Model(&role).Association("Menus").Delete(role.Menus)
	//最后删除角色
	err = global.DB.Where("id = ?", id).Delete(&model.Role{}).Error
	return err
}

// 更新用户关联的角色组 by 角色名数组
func UpdateUserRoleGroup(roles []string, user *model.User) error {
	var roleArr []model.Role
	global.DB.Where("role_name in ?", roles).Find(&roleArr)
	return global.DB.Debug().Model(&user).Association("RoleGroup").Replace(roleArr)
}

// 删除用户关联的角色组
func DeleteUserRoleGroup(user *model.User) error {
	return global.DB.Debug().Model(&model.User{ID: user.ID}).Association("RoleGroup").Replace(nil)
}
