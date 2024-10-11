package logic

import "fisco-cert-app/dao/fisco"

func GetBlockHeight() int64 {
	return fisco.GetBlockHeight()
}
