package model

import (
	"fmt"
	"goBlog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null;" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// CreateCategory 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCategory 查询分类列表
func GetCategory(pageSize, pageNum int) ([]Category, int64) {
	var category []Category
	var CurPageSize int
	var CurOffSet int
	var total int64
	//判断输入的pageSize是否为0,为0表示解除limit的限制，也就是设置-1
	if pageSize > 0 {
		CurPageSize = pageSize
	} else {
		CurPageSize = -1
	}
	//判断输入的pageNum是否为0,为0表示解除limit的限制，也就是设置-1
	if pageNum > 0 {
		CurOffSet = (CurOffSet - 1) * CurPageSize
	} else {
		CurOffSet = -1
	}
	err = db.Limit(CurPageSize).Offset(CurOffSet).Find(&category).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return category, total
}

// DeleteCategory 删除分类
func DeleteCategory(id int) (code int) {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		fmt.Println("test流程")
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditCategory 更新分类
func EditCategory(id int, data *Category) (code int) {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
