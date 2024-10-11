package mysql

import "fisco-cert-app/models"

func InsertAuth(auth *models.Auth) (err error) {
	result := db.Model(&models.Auth{}).Create(auth)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
