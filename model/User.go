package model

import (
	"fmt"
	"goBlog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(100);not null;" json:"password" validate:"required,min=6,max=20"`
	Role     int    `gorm:"type:varchar(20);DEFAULT:2;" json:"role" validate:"required,gte=1"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// CheckUpUser 允许修改用户的时候不修改名字, 并且如果修改的名字已经存在报错
func CheckUpUser(id int, name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = BcryptPassword([]byte(data.Password))
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetSingleUser 查询单个用户
func GetSingleUser(id int) (User, int) {
	var user User
	err := db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(userName string, pageSize, pageNum int) ([]User, int64) {
	var users []User
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
	if userName == "" {
		err = db.Limit(CurPageSize).Offset(CurOffSet).Find(&users).Offset(-1).Limit(-1).Count(&total).Error
	} else {
		err = db.Where("username LIKE ?", userName+"%").Limit(CurPageSize).Offset(CurOffSet).Find(&users).Offset(-1).Limit(-1).Count(&total).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// BcryptPassword 密码加密
func BcryptPassword(data []byte) string {
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.MinCost)
	if err != nil {
		log.Fatal("加密失败")
		return string(data)
	}
	return string(hash)
}

// ComparePasswords 验证密码
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

// DeleteUser 删除用户
func DeleteUser(id int) (code int) {
	var users User
	err = db.Where("id = ?", id).Delete(&users).Error
	if err != nil {
		fmt.Println("test流程")
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 更新用户
func EditUser(id int, data *User) (code int) {
	var users User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&users).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// CheckLogin 登录验证
func CheckLogin(username, password string) int {
	var users User
	db.Where("username = ?", username).First(&users)
	fmt.Println(users.Password)
	fmt.Println(BcryptPassword([]byte(password)))
	fmt.Println(password)
	if users.ID == 0 {
		return errmsg.ErrorUsernameNotExist
	}
	if !ComparePasswords(users.Password, password) {
		return errmsg.ErrorPasswordWrong
	}
	if users.Role != 2 {
		return errmsg.ErrorUserNoRight
	}

	return errmsg.SUCCESS
}
