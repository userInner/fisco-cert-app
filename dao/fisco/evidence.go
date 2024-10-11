package fisco

import (
	"database/sql"
	"fisco-cert-app/dao/mysql"
	"fisco-cert-app/models"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

// 版权证书

// 保存图片版权信息
func SaveImgInfo(hexPrivateKey, evi string, userID, imgageID uint64) (err error) {
	// 版权号
	zap.L().Info("SaveImgInfo params...",
		zap.String("hexPrivateKey", hexPrivateKey),
		zap.String("evi", evi),
		zap.Uint64("userID", userID),
		zap.Uint64("imgageID", imgageID))

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		zap.L().Debug("crypto.HexToECDSA auth failed", zap.Error(err))
		return err
	}

	u := bind.NewKeyedTransactor(privateKey)
	address, tx, _, err := evidenceFactory.NewEvidence(u, evi)

	cert := &models.Cert{
		TxID:     tx.Hash().Hex(),
		UserID:   userID,
		ImageID:  imgageID,
		Evidence: address.Hex(),
	}

	err = mysql.InsertCert(cert)
	if err != nil {
		zap.L().Debug("mysql isnertCert failed", zap.Error(err))
		return err
	}

	// 记录用户上链记录
	// 查询是否有记录,有则加，无则建
	monitor, err := mysql.GetUserTxMonitor(userID)
	if err != nil {
		return err
	}
	monitor.TransCount += 1
	monitor.TransHashLastest = sql.NullString{tx.Hash().Hex(), true}
	monitor.UserID = userID
	return mysql.UpdateTxMonitor(monitor)
}
