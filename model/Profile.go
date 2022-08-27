package model

import "goBlog/utils/errmsg"

type Profile struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20)" json:"name"`
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	Email  string `gorm:"type:varchar(200)" json:"email"`
	Img    string `gorm:"type:varchar(300)" json:"img"`
	Avatar string `gorm:"type:varchar(300)" json:"avatar"`
}

// GetProfile 获得个人信息配置
func GetProfile(id int) (Profile, int) {
	var profile Profile
	err := db.Where("id = ?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

// UpdateProfile 更新个人信息
func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	err := db.Model(&profile).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
