package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 基于中间件获取的UserID
func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	// 未获取到值
	if !ok {
		ResponseError(c, CodeNeedLogin)
		return
	}
	userID, ok = uid.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// getAccessToken 获取jwt中生成的新的AccessToken
func getAccessToken(c *gin.Context) (accessToken string, err error) {
	AToken, ok := c.Get(CtxUserAccessTokenKey)
	if !ok {
		ResponseError(c, CodeNeedLogin)
		return "", ErrorUserNotLogin
	}
	accessToken, ok = AToken.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// getUserAddr 获取中间件携带的用户地址
func getUserAddr(c *gin.Context) (address string, err error) {
	addr, ok := c.Get(CtxUserAddrKey)
	if !ok {
		ResponseError(c, CodeNeedLogin)
		return "", ErrorUserNotLogin
	}
	address, ok = addr.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// getPageInfo 获取请求中的页码
func getPageInfo(c *gin.Context) (page, size int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var err error
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return
}
