package mysql

import "fisco-cert-app/models"

func InsertCert(cert *models.Cert) (err error) {
	result := db.Model(&models.Cert{}).Create(cert)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCertDetail(txID string) (cert *models.Cert, err error) {
	err = db.Model(&models.Cert{}).
		Select("tx_id,user_id,image_id,evidence,created_at").
		Where("tx_id", txID).
		First(&cert).Error
	return
}

func GetUserCertSum(uid uint64) (sum int64, err error) {
	err = db.Model(&models.Cert{}).
		Count(&sum).
		Where("user_id", uid).Error
	return
}

func GetCertListByUser(userID uint64) (certs []*models.Cert, err error) {
	err = db.Model(&models.Cert{}).
		Select("tx_id,user_id,image_id,evidence,created_at").
		Where("user_id", userID).
		//Limit(int(p.Size)).
		//Offset(int((p.Page - 1) * p.Size)).
		Find(&certs).Error
	return
}

func GetCertList(p *models.ParamEvidenceList) (certs []*models.Cert, err error) {
	err = db.Model(&models.Cert{}).
		Select("tx_id,user_id,image_id,evidence,created_at").
		Limit(int(p.Size)).
		Offset(int((p.Page - 1) * p.Size)).
		Find(&certs).Error
	return
}
