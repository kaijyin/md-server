package system

import (
	"errors"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/table"
	"github.com/kaijyin/md-server/server/utils"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.User
//@return: err error, userInter model.User

type UserService struct {
}

func (userService *UserService) Register(u table.User) (err error, userInter table.User) {
	var user table.User
	if !errors.Is(global.MD_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.MD_DB.Create(&u).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.User
//@return: err error, userInter *model.User

func (userService *UserService) Login(u *table.User) (err error, userInter *table.User) {
	var user table.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.MD_DB.Where("username = ? ", u.Username).First(&user).Error
	if errors.Is(gorm.ErrRecordNotFound,err){
		//不存在就新建
		err=global.MD_DB.Create(&user).Error
	}else{
		err = global.MD_DB.Where("username = ? and password = ?", u.Username, u.Password).First(&user).Error
	}
	return err, &user
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.User, newPassword string
//@return: err error, userInter *model.User

func (userService *UserService) ChangePassword(u *table.User, newPassword string) (err error, userInter *table.User) {
	var user table.User
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.MD_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.MD_DB.Model(&table.User{})
	var userList []table.User
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return err, userList, total
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.User
//@return: err error, user model.User

func (userService *UserService) SetUserInfo(reqUser table.User) (err error, user table.User) {
	err = global.MD_DB.Updates(&reqUser).Error
	return err, reqUser
}


//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.User

func (userService *UserService) FindUserById(id uint) (err error, user *table.User) {
	var u table.User
	err = global.MD_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

func (userService *UserService) DeleteUser(id uint) (err error) {
	err = global.MD_DB.Delete(&table.User{},id).Error
	return err
}
