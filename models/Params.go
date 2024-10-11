package models

import "mime/multipart"

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`                     // 用户名
	Password   string `json:"password" binding:"required"`                     // 用户密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 用户再次输入密码
}

// ParamLogin 登陆请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 用户密码
}

// ParamEvidence 创建证书请求参数
type ParamEvidence struct {
	HexPrivateKey string                `form:"hexPrivateKey" json:"hexPrivateKey" binding:"required"` // 私钥
	Evidence      string                `form:"evidence" json:"evidence" binding:"required"`           // 存证信息
	File          *multipart.FileHeader `form:"file" type:"blob" binding:"required"`                   // 添加水印的图片
}

// ParamImgEvidence 图片水印存证
type ParamImgEvidence struct {
	WatermarkURL  string `form:"watermark" json:"watermark" binding:"required"`         // 水印图片
	OriginImg     string `form:"img" json:"img" binding:"required"`                     // 原图
	HexPrivateKey string `form:"hexPrivateKey" json:"hexPrivateKey" binding:"required"` // 私钥
}

// ParamEvidence 创建证书请求参数
type ParamAuth struct {
	File    *multipart.FileHeader `form:"file" type:"blob" binding:"required"`
	Address string                `form:"address" json:"address" binding:"required"`
}

// ParamEvidenceList 获取版权列表query string参数
type ParamEvidenceList struct {
	Page     int64  `json:"page" form:"page" example:"1"`          // 页码
	Size     int64  `json:"size" form:"size" example:"10"`         // 每页数据量
	UserName string `json:"user_name" form:"user_name" example:""` // 用户名
}

// ParamEvidenceList 获取版权列表query string参数
type ParamEvidenceDeatil struct {
	TxID string `json:"tx_id" form:"tx_id" binding:"required"`
}

// ParamWaterMark 图片水印
type ParamWaterMark struct {
	File *multipart.FileHeader `form:"file" type:"blob" binding:"required"`
	Type uint8                 `form:"type" json:"type" binding:"required" example:"1"`
}

// ParamPage
type Page struct {
	Total int `json:"total"` // 总页数
}
