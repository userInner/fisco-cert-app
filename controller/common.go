package controller

import (
	"fisco-cert-app/dao/mysql"
	"fisco-cert-app/logic"
	"fisco-cert-app/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// WaterMarkHandler 上传需要添加水印的图片
// @Summary 添加文字水印
// @Description 通过用户私钥，为原图添加盲水印
// @Tags 水印相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamEvidence true "查询参数"
// @Param Authorization header string false "Bearer JWT"
// @Success 200 {object} string "url"
// @Router /watermark [post]
func WaterMarkHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)

	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	evidence := new(models.ParamEvidence)
	if err = c.ShouldBind(evidence); err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("invalid params", zap.Error(err))
		return
	}

	// 查询用户ID
	user, err := mysql.GetUserByID(userID)
	if err != nil {
		ResponseError(c, CodeInvalidToken)
		zap.L().Info("mysql.GetUserByID failed", zap.Error(err))
		return
	}

	url, err := logic.UploadFile(userID, user.Address, evidence)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Info("logic.UploadFile failed", zap.Error(err))
		return
	}

	ResponseSuccess(c, url)
}

// WaterMarkImgHandler
// @Summary 添加图片水印
// @Description 通过用户私钥，为原图添加盲水印
// @Tags 水印相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamImgEvidence true "查询参数"
// @Param Authorization header string false "Bearer JWT"
// @Success 200 {object} string "url"
// @Router /img [post]
func WaterMarkImgHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)

	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	evidence := new(models.ParamImgEvidence)
	if err = c.ShouldBindJSON(evidence); err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("invalid params", zap.Error(err))
		return
	}

	// 查询用户ID
	user, err := mysql.GetUserByID(userID)
	if err != nil {
		ResponseError(c, CodeInvalidToken)
		zap.L().Info("mysql.GetUserByID failed", zap.Error(err))
		return
	}

	// 生成盲水印
	url, err := logic.GetImgWaterMark(userID, user.Address, evidence)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Info("logic.GetImgWaterMark failed", zap.Error(err))
		return
	}

	ResponseSuccess(c, url)
}

// AuthHandler 文字水印图片鉴权
// @Summary 对图片盲水印进行提取
// @Description 对原图片进行盲水印提取
// @Tags 水印相关接口
// @Accept form-data/form-data
// @Produce application/json
// @Param object body models.ParamAuth true "查询参数"
// @Param Authorization header string false "Bearer JWT"
// @Success 200 {object} string "url"
// @Router /auth [post]
func AuthHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	p := new(models.ParamAuth)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 拿到水印图
	url, err := logic.GetWaterMark(userID, p)
	if err != nil {
		zap.L().Error("logic.GetWaterMark failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err)
		return
	}

	ResponseSuccess(c, url)
}

// ImgAuthHandler 图片水印鉴权
// @Summary 图片水印提取
// @Description 对原图片进行盲水印提取
// @Tags 水印相关接口
// @Accept form-data/form-data
// @Produce application/json
// @Param object body models.ParamImgEvidence true "查询参数"
// @Param Authorization header string false "Bearer JWT"
// @Success 200 {object} string "url"
// @Router /img [post]
func ImgAuthHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	evidence := new(models.ParamImgEvidence)
	if err = c.ShouldBindJSON(evidence); err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("invalid params", zap.Error(err))
		return
	}

	// 查询用户ID
	user, err := mysql.GetUserByID(userID)
	if err != nil {
		ResponseError(c, CodeInvalidToken)
		zap.L().Info("mysql.GetUserByID failed", zap.Error(err))
		return
	}
	// 生成盲水印
	url, err := logic.ImgAuth(userID, user.Address, evidence)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Info("logic.GetImgWaterMark failed", zap.Error(err))
		return
	}

	ResponseSuccess(c, url)
}

// GetEvidenceListHandler 查询版权列表
// @Summary 查询版权列表
// @Description 查询上链的版权信息
// @Tags 水印相关接口
// @Produce application/json
// @Param object query models.ParamEvidenceList true "查询参数"
// @Param Authorization header string false "Bearer JWT"
// @Success 200 {object} []models.EvidenceResp "data"
// @Router /list [get]
func GetEvidenceListHandler(c *gin.Context) {
	p := new(models.ParamEvidenceList)
	if err := c.ShouldBindQuery(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("controller.GetEvidenceList failed", zap.Error(err))
		return
	}
	evidences, err := logic.GetEvidenceList(p)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("logic.GetEvidenceList failed", zap.Error(err))
		return
	}

	ResponseSuccess(c, evidences)
}

// GetEvidenceDetailHandler 查看版权详情
// @Summary 查看版权详情
// @Description 查询上链的版权详情信息
// @Tags 水印相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamEvidenceDeatil true "查询参数"
// @Param Authorization header string false "Bearer JWT"
// @Success 200 {object} models.EvidenceDetailResp "data"
// @Router /evidence [get]
func GetEvidenceDetailHandler(c *gin.Context) {
	p := new(models.ParamEvidenceDeatil)
	if err := c.ShouldBindQuery(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("controller.GetEvidenceDetailHandler failed", zap.Error(err))
		return
	}

	evidence, err := logic.GetEvidenceDetail(p)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("controller.GetEvidenceDetail failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, evidence)
}

// UpWaterMarkHandler 上传水印
// @Summary 上传图片水印
// @Description 上传水印图片，为后续图片水印进行处理
// @Tags 水印相关接口
// @Accept form-data/form-data
// @Produce application/json
// @Param object body models.ParamWaterMark true "查询参数"
// @Param Authorization header string false "Bearer JWT"
// @Success 200 {object} string "url"
// @Router /upwatermark [post]
func UpWaterMarkHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	waterMark := new(models.ParamWaterMark)
	if err = c.ShouldBind(waterMark); err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("invalid params", zap.Error(err))
		return
	}

	// 查询用户ID
	user, err := mysql.GetUserByID(userID)
	if err != nil {
		ResponseError(c, CodeInvalidToken)
		zap.L().Info("mysql.GetUserByID failed", zap.Error(err))
		return
	}

	url, err := logic.UpWaterMark(user.UserID, waterMark)
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err)
		zap.L().Info("logic.UpWaterMark failed", zap.Error(err))
		return
	}
	// 返回上传的链接
	ResponseSuccess(c, url)
}

// GetUserCountHandler 用户数量
// @Summary 用户数量
// @Description 用户数量
// @Tags 水印相关接口
// @Produce application/json
// @Success 200 {object} int64 "data"
// @Router /getUserCount [get]
func GetUserCountHandler(c *gin.Context) {
	userCount, err := mysql.GetUserCount()
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err)
		zap.L().Error("controller.GetEvidenceDetail failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, userCount)
}
