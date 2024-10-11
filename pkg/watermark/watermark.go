package watermark

import (
	"context"
	"encoding/base64"
	"errors"
	"fisco-cert-app/setting"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var errWaterMarkInitFailed = errors.New("水印初始化失败")

const (
	up               = "watermark/3/type/3/text/" // 上传
	down             = "watermark/4/type/3/text"  // 提取
	Origin           = "-origin"                  // 原图后缀
	Marked           = "-marked"                  // 添加盲水印后的后缀
	Mark             = "-mark"                    // 水印内容后缀
	suffix           = "/version/2.0"             // 使用新版盲水印算法
	presignedURLTime = time.Hour * 24 * 7         // 七天更新
)

const (
	upByImg   = "watermark/3/type/2/image/" // 上传
	downByImg = "watermark/4/type/2/image/" // 提取
)

var (
	tencentConfig *setting.Tencent
)

// Init 初始化配置
func Init(tencent *setting.Tencent) {
	tencentConfig = tencent
}

func NewClient() *cos.Client {
	u, _ := url.Parse(tencentConfig.CosRegion)
	//// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse(tencentConfig.GetService)
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  tencentConfig.SecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: tencentConfig.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})

	client.Conf.EnableCRC = false
	return client

}

// UploadImage 上传水印图片
func UploadWaterMark(addr, name, imgSuffix string, reader io.Reader) (err error) {
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			XOptionHeader: &http.Header{},
		},
	}

	pic := &cos.PicOperations{
		IsPicInfo: 1,
		Rules: []cos.PicOperationsRules{
			{
				FileId: name + Marked + imgSuffix,
				Rule:   upByImg + urlSafeBase64Encode(addr) + suffix,
			},
		},
	}

	opt.XOptionHeader.Add("Pic-Operations", cos.EncodePicOperations(pic))
	// 放置对象存储进行持久化存储
	_, _, err = NewClient().CI.Put(context.Background(), name+Origin+imgSuffix, reader, opt)
	if err != nil {
		zap.L().Error("tencent watermark.UploadImage function failed", zap.Error(err))
	}
	return
}

// UploadImage 生成盲水印图片
func UploadImage(url, name, imgSuffix string, reader io.Reader) (err error) {
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			XOptionHeader: &http.Header{},
		},
	}
	pic := &cos.PicOperations{
		IsPicInfo: 1,
		Rules: []cos.PicOperationsRules{
			{
				FileId: name + Marked + imgSuffix,
				Rule:   up + urlSafeBase64Encode(url) + suffix,
			},
		},
	}

	opt.XOptionHeader.Add("Pic-Operations", cos.EncodePicOperations(pic))
	// 放置对象存储进行持久化存储
	_, _, err = NewClient().CI.Put(context.Background(), name+Origin+imgSuffix, reader, opt)
	if err != nil {
		zap.L().Error("tencent watermark.UploadImage function failed", zap.Error(err))
	}
	return
}

// GetWaterMark 提取已提交盲水印成品的水印图
func GetWaterMark(name, imgSuffix string) (err error) {
	opt := &cos.ImageProcessOptions{
		IsPicInfo: 1,
		Rules: []cos.PicOperationsRules{
			{
				FileId: name + Marked + imgSuffix,
				Rule:   down + suffix,
			},
		},
	}

	// 对桶内文件提取
	_, _, err = NewClient().CI.ImageProcess(context.Background(), name+Marked+imgSuffix, opt)
	if err != nil {
		zap.L().Error("tencent watermark.GetWaterMark function failed", zap.Error(err))
	}
	return
}

// GetImgWaterMark 生成含水印图(图片水印)
func GetImgWaterMark(uid, imgSuffix, watermark, origin string) (err error) {
	zap.L().Info("uid, imgSuffix, watermark, origin", zap.Strings("...", []string{uid, imgSuffix, watermark, origin}))
	opt := &cos.ImageProcessOptions{
		IsPicInfo: 1,
		Rules: []cos.PicOperationsRules{
			{
				FileId: uid + Marked + imgSuffix,
				Rule:   upByImg + urlSafeBase64Encode(watermark) + suffix,
			},
		},
	}

	// 对桶内文件提取
	_, _, err = NewClient().CI.ImageProcess(context.Background(), origin, opt)
	if err != nil {
		zap.L().Error("tencent watermark.GetWaterMark function failed", zap.Error(err))
	}
	return
}

func urlSafeBase64Encode(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "/", "_")
	encoded = strings.ReplaceAll(encoded, "=", "")
	return encoded
}

// GetGenImgWaterMark 鉴别图片获得水印(图片水印)
func GetGenImgWaterMark(uid, imgSuffix, watermark, origin string) (err error) {
	zap.L().Info("params:", zap.Strings("...", []string{uid, imgSuffix, watermark, origin}))
	opt := &cos.ImageProcessOptions{
		IsPicInfo: 1,
		Rules: []cos.PicOperationsRules{
			{
				FileId: uid + Mark + imgSuffix,
				Rule:   downByImg + urlSafeBase64Encode(watermark) + suffix,
			},
		},
	}
	fmt.Println(3)
	// 对桶内文件提取
	_, _, err = NewClient().CI.ImageProcess(context.Background(), origin, opt)
	if err != nil {
		zap.L().Error("tencent watermark.GetWaterMark function failed", zap.Error(err))
	}
	return
}

// GetUserImageWatermark 对于用户提交的盲水印图片进行提取
func GetUserImageWatermark(name, imgSuffix string, reader io.Reader) (err error) {
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			XOptionHeader: &http.Header{},
		},
	}

	fileid := name + Mark + imgSuffix
	pic := &cos.PicOperations{
		IsPicInfo: 1,
		Rules: []cos.PicOperationsRules{
			{
				FileId: fileid, // 把用户上传的需要提取盲水印图片认为已添加的盲水印图片
				Rule:   down + "/" + urlSafeBase64Encode(fileid) + suffix,
				//Rule: down + suffix,
			},
		},
	}

	opt.XOptionHeader.Add("Pic-Operations", cos.EncodePicOperations(pic))
	// 放置对象存储进行持久化存储
	_, _, err = NewClient().CI.Put(context.Background(), fileid, reader, opt)
	if err != nil {
		zap.L().Error("tencent watermark.GetUserImageWatermark function failed", zap.Error(err))
	}
	return
}

// UploadFile 上传普通文件
func UploadFile(name string, reader io.Reader) (err error) {
	_, err = NewClient().Object.Put(context.Background(), name, reader, nil)
	if err != nil {
		zap.L().Error("tencent watermark.UploadFile function failed", zap.Error(err))
	}
	return
}

// DeleteFile 删除某个文件
func DeleteFile(name string) (err error) {
	_, err = NewClient().Object.Delete(context.Background(), name, nil)
	if err != nil {
		zap.L().Error("tencent watermark.DeleteFile function failed", zap.Error(err))
	}
	return
}

// GetPresignedURL cos生成临时链接
func GetPresignedURL(name string) (string, error) {
	presignedURL, err := NewClient().Object.GetPresignedURL(context.Background(),
		http.MethodGet,
		name,
		tencentConfig.SecretId,
		tencentConfig.SecretKey,
		presignedURLTime,
		nil)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
