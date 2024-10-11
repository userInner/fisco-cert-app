package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"fisco-cert-app/models"
	"fisco-cert-app/setting"
)

func CheckUserExist(u *models.User) (err error) {
	var count int64
	result := db.Model(&models.User{}).
		Where("user_name", u.UserName).
		Count(&count)
	if result.Error != nil { // gorm查询系统错误
		return result.Error
	}
	if count > 0 { // 用户已存在
		return ErrorUserExist
	}
	return nil
}

func InsertUser(u *models.User) (err error) {
	u.Password = md5Pwd(u.Password)
	result := db.Model(&models.User{}).Create(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByUserName(userName string) (users []*models.User, err error) {
	result := db.Model(&models.User{}).
		Where("user_name like ?", "%"+userName+"%").
		Select("user_id, user_name").
		Find(&users)
	if result.Error != nil { // gorm查询系统错误
		return nil, result.Error
	}
	return
}

func GetUserCount() (count int64, err error) {
	result := db.Model(&models.User{}).
		Count(&count).Error
	if result != nil { // gorm查询系统错误
		return 0, result
	}
	return
}

func GetUserByID(userID uint64) (user *models.User, err error) {
	err = db.Model(&models.User{}).
		Select("user_id, user_name, address").
		Where("user_id", userID).
		First(&user).Error
	return
}

func Login(u *models.User) (error error) {
	oldMd5Pwd := u.Password
	result := db.Model(&models.User{}).
		Select("user_id, user_name, password").
		Where("user_name = ?", u.UserName).
		Find(u)
	if result.Error != nil {
		return result.Error
	}
	if md5Pwd(oldMd5Pwd) != u.Password {
		return ErrorInvalidPassword
	}
	return
}

func md5Pwd(password string) string {
	h := md5.New()
	h.Write([]byte(setting.Conf.AuthConfig.PwdSecret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
