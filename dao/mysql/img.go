package mysql

import "fisco-cert-app/models"

func InsertImg(img *models.Img) (err error) {
	result := db.Model(&models.Img{}).Create(img)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCertByID(imgID uint64) (img *models.Img, err error) {
	err = db.Model(&models.Img{}).
		Select("img_id, file_name, suffix").
		Where("img_id", imgID).
		First(&img).Error
	return
}
