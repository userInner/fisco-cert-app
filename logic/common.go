package logic

import (
	"fisco-cert-app/dao/fisco"
	"fisco-cert-app/dao/mysql"
	"fisco-cert-app/dao/redis"
	"fisco-cert-app/models"
	"fisco-cert-app/pkg/snowflake"
	"fisco-cert-app/pkg/watermark"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"path"
	"strings"
)

/**
通用接口
上传
*/

// UploadFile 用户上传文件后，返回已经添加盲水印的图片
func UploadFile(userID uint64, addr string, evidence *models.ParamEvidence) (url string, err error) {
	// 比对address 和 hexPrivateKey
	flag, err := fisco.EqualsHexPrivateKeyAddress(evidence.HexPrivateKey, addr)
	if !flag {
		return "", err
	}

	f, err := evidence.File.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	name := uuid.New().String()
	suffix := path.Ext(evidence.File.Filename)

	// 添加盲水印
	err = watermark.UploadImage(addr, name, suffix, f)
	if err != nil {
		return
	}

	// 返回临时盲水印图链接
	url, err = watermark.GetPresignedURL(name + watermark.Marked + suffix)
	if err != nil {
		return
	}
	// 存储水印在桶内的名字
	imgID := uint64(snowflake.GenID())

	img := &models.Img{
		ImgID:    imgID,
		UserID:   userID,
		FileName: name + watermark.Marked,
		FileSize: uint64(evidence.File.Size), // 原始图大小
		Suffix:   suffix,
	}
	err = mysql.InsertImg(img)
	if err != nil {
		return
	}
	// 存储临时盲水印图片链接
	err = redis.SavePresignedURL(name+watermark.Marked, url)
	if err != nil {
		return
	}
	// 版权上链
	err = fisco.SaveImgInfo(evidence.HexPrivateKey, evidence.Evidence, userID, imgID)
	if err != nil {
		return
	}

	return
}

// ImgAuthHandler 鉴权
func ImgAuth(userId uint64, addr string, p *models.ParamImgEvidence) (url string, err error) {
	// 比对登陆用户与上传私钥是否相同
	flag, err := fisco.EqualsHexPrivateKeyAddress(p.HexPrivateKey, addr)
	if !flag {
		return "", err
	}

	// 获取redis上传鉴权图像，以及水印图
	waterMarkURL, err := redis.GetPresignedURL(strings.Split(p.WatermarkURL, ".")[0])
	if err != nil {
		zap.L().Error("redis GetPresignedURL not exist", zap.Error(err))
		return "", err
	}
	uid := uuid.New().String()

	// 获取水印
	suffix := strings.Split(p.WatermarkURL, ".")[1]
	httpLink := strings.Replace(waterMarkURL, "https://", "http://", 1)
	err = watermark.GetGenImgWaterMark(uid, "."+suffix, httpLink, p.OriginImg)
	if err != nil {
		zap.L().Error("tencent GetImgWaterMark failed", zap.Error(err))
		return "", err
	}
	fmt.Println(2)

	// 获取该用户提交的图片水印
	log.Info("url name", zap.String("url", uid+watermark.Mark+suffix))
	url, err = watermark.GetPresignedURL(uid + watermark.Mark + "." + suffix)
	if err != nil {
		return "", err
	}
	imgID := uint64(snowflake.GenID())

	// 上传img表
	img := &models.Img{
		ImgID:    imgID,
		UserID:   userId,
		FileName: uid + watermark.Mark,
		FileSize: 0,
		Suffix:   "." + suffix,
	}
	if err = mysql.InsertImg(img); err != nil {
		return "", err
	}

	// 链下记录存证hash
	auth := &models.Auth{
		AuthID:  uint64(snowflake.GenID()),
		UserID:  userId,
		ImageID: imgID,
		Address: addr,
	}
	if err = mysql.InsertAuth(auth); err != nil {
		return "", err
	}

	// redis记录url
	if err = redis.SavePresignedURL(uid+watermark.Marked, url); err != nil {
		return "", err
	}
	return
}

// GetWaterMark 获取用户上传的图片水印
func GetWaterMark(userId uint64, p *models.ParamAuth) (url string, err error) {
	f, err := p.File.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	uid := uuid.New().String()
	suffix := path.Ext(p.File.Filename)

	if err = watermark.GetUserImageWatermark(uid, suffix, f); err != nil {
		return
	}
	// 获取该用户提交的图片水印
	url, err = watermark.GetPresignedURL(uid + watermark.Mark + suffix)
	if err != nil {
		return "", err
	}
	imgID := uint64(snowflake.GenID())
	// 上传img表
	img := &models.Img{
		ImgID:    imgID,
		UserID:   userId,
		FileName: uid + watermark.Marked,
		FileSize: uint64(p.File.Size),
		Suffix:   suffix,
	}

	if err = mysql.InsertImg(img); err != nil {
		return "", err
	}

	auth := &models.Auth{
		AuthID:  uint64(snowflake.GenID()),
		UserID:  userId,
		ImageID: imgID,
		Address: p.Address,
	}

	if err = mysql.InsertAuth(auth); err != nil {
		return "", err
	}
	if err = redis.SavePresignedURL(uid+watermark.Marked, url); err != nil {
		return "", err
	}
	return
}

// GetImgWaterMark 获取用户上传的图片水印
func GetImgWaterMark(userId uint64, addr string, p *models.ParamImgEvidence) (url string, err error) {
	// 比对登陆用户与上传私钥是否相同
	flag, err := fisco.EqualsHexPrivateKeyAddress(p.HexPrivateKey, addr)
	if !flag {
		return "", err
	}

	// 获取redis中的url
	waterMarkURL, err := redis.GetPresignedURL(strings.Split(p.WatermarkURL, ".")[0])
	if err != nil {
		zap.L().Error("redis GetPresignedURL not exist", zap.Error(err))
		return "", err
	}
	uid := uuid.New().String()
	// 添加水印
	suffix := strings.Split(p.WatermarkURL, ".")[1]
	httpLink := strings.Replace(waterMarkURL, "https://", "http://", 1)
	err = watermark.GetImgWaterMark(uid, "."+suffix, httpLink, p.OriginImg)
	if err != nil {
		zap.L().Error("tencent GetImgWaterMark failed", zap.Error(err))
		return "", err
	}

	// 获取该用户提交的图片水印
	log.Info("url name", zap.String("url", uid+watermark.Marked+suffix))
	url, err = watermark.GetPresignedURL(uid + watermark.Marked + "." + suffix)
	if err != nil {
		return "", err
	}
	imgID := uint64(snowflake.GenID())

	// 上传img表
	img := &models.Img{
		ImgID:    imgID,
		UserID:   userId,
		FileName: uid + watermark.Marked,
		FileSize: 0,
		Suffix:   suffix,
	}

	if err = mysql.InsertImg(img); err != nil {
		return "", err
	}

	// 存放到区块链中
	err = fisco.SaveImgInfo(p.HexPrivateKey, uid+watermark.Marked+suffix, userId, imgID)
	if err != nil {
		return "", err
	}

	// 链下记录存证hash
	auth := &models.Auth{
		AuthID:  uint64(snowflake.GenID()),
		UserID:  userId,
		ImageID: imgID,
		Address: addr,
	}

	if err = mysql.InsertAuth(auth); err != nil {
		return "", err
	}

	// redis记录url
	if err = redis.SavePresignedURL(uid+watermark.Marked, url); err != nil {
		return "", err
	}
	return
}

// UpWaterMark 用户上传的图片水印
func UpWaterMark(userId uint64, p *models.ParamWaterMark) (url string, err error) {
	f, err := p.File.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	uid := uuid.New().String()
	suffix := path.Ext(p.File.Filename)
	// 判断上传原图还是水印
	quality := watermark.Mark
	if p.Type == 2 {
		quality = watermark.Origin
	}

	if err = watermark.UploadFile(uid+quality+suffix, f); err != nil {
		return
	}
	// 获取该用户提交的图片水印
	url, err = watermark.GetPresignedURL(uid + quality + suffix)
	if err != nil {
		return "", err
	}
	imgID := uint64(snowflake.GenID())
	// 上传img表
	img := &models.Img{
		ImgID:    imgID,
		UserID:   userId,
		FileName: uid + quality,
		FileSize: uint64(p.File.Size),
		Suffix:   suffix,
	}

	if err = mysql.InsertImg(img); err != nil {
		return "", err
	}
	if err = redis.SavePresignedURL(uid+quality, url); err != nil {
		return "", err
	}
	// 返回图片ID
	return uid + quality + suffix, nil
}

// GetEvidenceList 查询版权列表
func GetEvidenceList(p *models.ParamEvidenceList) (evidences []*models.EvidenceResp, err error) {
	// 基于cert表查询版权信息
	certs := make([]*models.Cert, 0)
	zap.L().Info("GetEvidenceList param user_name", zap.String("user_name", p.UserName))
	if p.UserName == "" {
		certs, err = mysql.GetCertList(p)
		if err != nil {
			return nil, err
		}
	} else {
		users, err := mysql.GetUserByUserName(p.UserName)
		if err != nil {
			return nil, err
		}
		for _, item := range users {
			userCerts, err := mysql.GetCertListByUser(item.UserID)
			if err != nil {
				zap.L().Error("getCertListByUser failed,", zap.Error(err))
				continue
			}
			certs = append(certs, userCerts...)
		}
	}

	evidences = make([]*models.EvidenceResp, len(certs))
	// 查询版权用户信息
	for key, v := range certs {
		// 查询版权用户信息
		user, err := mysql.GetUserByID(v.UserID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID failed", zap.Uint64("userID", v.UserID), zap.Error(err))
			continue
		}

		// 查询版权图片信息
		img, err := mysql.GetCertByID(v.ImageID)
		if err != nil {
			zap.L().Error("mysql.GetCertByID failed", zap.Uint64("imageID", v.ImageID), zap.Error(err))
			continue
		}

		// 基于redis 查询 url
		url, err := redis.GetPresignedURL(img.FileName)
		if err != nil {
			zap.L().Error("GetPresignedURL failed", zap.String("name", img.FileName+watermark.Marked), zap.Error(err))
			continue
		}

		evi := &models.EvidenceResp{
			TxID:      v.TxID,
			ImgURL:    url,
			UserName:  user.UserName,
			CreatedAt: v.CreatedAt,
		}

		evidences[key] = evi
	}
	return evidences, nil
}

// GetEvidenceDetail 查询版权列表
func GetEvidenceDetail(p *models.ParamEvidenceDeatil) (evidence *models.EvidenceDetailResp, err error) {
	// 基于cert表查询版权信息
	cert, err := mysql.GetCertDetail(p.TxID)
	if err != nil {
		zap.L().Error("mysql.GetCertDetail failed", zap.String("tx_id", cert.TxID), zap.Error(err))
		return nil, err
	}
	// 查询版权用户信息
	user, err := mysql.GetUserByID(cert.UserID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID failed", zap.Uint64("userID", cert.UserID), zap.Error(err))
		return nil, err
	}

	// 查询版权图片信息
	img, err := mysql.GetCertByID(cert.ImageID)
	if err != nil {
		zap.L().Error("mysql.GetCertByID failed", zap.Uint64("imageID", cert.ImageID), zap.Error(err))
		return nil, err
	}

	// 基于redis 查询 url
	url, err := redis.GetPresignedURL(img.FileName)
	if err != nil {
		zap.L().Error("GetPresignedURL failed", zap.String("name", img.FileName+watermark.Marked+img.Suffix), zap.Error(err))
		return nil, err
	}

	evi := &models.EvidenceDetailResp{
		EvidenceResp: &models.EvidenceResp{
			TxID:     cert.TxID,
			ImgURL:   url,
			UserName: user.UserName,
		},
		Evidence:  cert.Evidence,
		CreatedAt: cert.CreatedAt,
	}
	return evi, nil
}
