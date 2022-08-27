package model

import (
	"fmt"
	"goBlog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(20);not null;" json:"title"'`
	Category Category `gorm:"foreignKey:Cid"`
	Cid      int      `gorm:"type: int;not null;" json:"cid"'`
	Desc     string   `gorm:"type: varchar(200);" json:"desc"` //表示describe
	Content  string   `gorm:"type:longtext;" json:"content""`
	Img      string   `gorm:"type:varchar(300);" json:"img"`
}

// CheckUser 查询用户是否存在
//func CheckUser(name string) (code int) {
//	var users User
//	db.Select("id").Where("username = ?", name).First(&users)
//	if users.ID > 0 {
//		return errmsg.ErrorUsernameUsed
//	}
//	return errmsg.SUCCESS
//}

// CreateArticle 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//GetAllArticle 查找分类下所有的文章
func GetAllArticle(id, pageSize, pageNum int) ([]Article, int, int64) {
	var article []Article
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
	err = db.Preload("Category").Limit(CurPageSize).Offset(CurOffSet).Where("cid = ?", id).Find(&article).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ErrorArticleNotExisted, 0
	}
	return article, errmsg.SUCCESS, total
}

//GetSingleArticle 查找单个文章
func GetSingleArticle(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ErrorArticleNotExisted
	}
	return article, errmsg.SUCCESS
}

// GetArticle 查找文章列表
func GetArticle(pageSize, pageNum int) ([]Article, int, int64) {
	var article []Article
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
	err = db.Preload("Category").Limit(CurPageSize).Offset(CurOffSet).Find(&article).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.SUCCESS, 0
	}
	return article, errmsg.SUCCESS, total
}

// DeleteArticle 删除文章
func DeleteArticle(id int) (code int) {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		fmt.Println("test流程")
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditArticle 更新文章
func EditArticle(id int, data *Article) (code int) {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
