package logic

import (
	"database/sql"
	"fisco-cert-app/dao/fisco"
	"fisco-cert-app/dao/mysql"
	"fisco-cert-app/models"
	"fisco-cert-app/pkg/jwt"
	"fisco-cert-app/pkg/snowflake"
)

func SignUp(u *models.ParamSignUp) (hexPrivateKey string, err error) {
	// 判断用户是否存在
	user := &models.User{}
	user.UserName = u.Username
	user.Password = u.Password
	if err = mysql.CheckUserExist(user); err != nil {
		return
	}
	// 生成snowflake ID
	uid := snowflake.GenID()
	user.UserID = uint64(uid)

	var address string
	// 生成非国密用户
	address, hexPrivateKey, err = fisco.GenUser()
	if err != nil {
		return
	}
	user.Address = address
	// 插入数据表
	if err = mysql.InsertUser(user); err != nil {
		return
	}

	monitor := &models.TransactionMonitor{
		UserID:           user.UserID,
		TransCount:       0,
		TransHashLastest: sql.NullString{"", false},
	}
	err = mysql.InsertTxMonitor(monitor)
	return
}

func Login(u *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{}
	user.UserName = u.Username
	user.Password = u.Password
	err = mysql.Login(user)
	if err != nil {
		return nil, err
	}
	AToken, RToken, err := jwt.GenBothToken(user.UserID, user.UserName)
	if err != nil {
		return nil, err
	}
	user.AccessToken = AToken
	user.RefreshToken = RToken
	return
}

func GetUserInfo(uid uint64) (user *models.User, err error) {
	user = &models.User{}
	user, err = mysql.GetUserByID(uid)
	if err != nil {
		return nil, err
	}
	// 查询存证数量
	certSum, err := mysql.GetUserCertSum(uid)
	user.CertSum = uint64(certSum)
	return
}
