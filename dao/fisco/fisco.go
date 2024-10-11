package fisco

import (
	"context"
	factory_contract "fisco-cert-app/dao/fisco/contract"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

var (
	fiscoClient     *client.Client
	evidenceFactory *factory_contract.EvidenceFactory
)

func Init(config *conf.Config) (err error) {
	fiscoClient, err = client.DialContext(context.Background(), config)
	if err != nil {
		return
	}
	addrs := make([]common.Address, 0, 1)
	addrs = append(addrs, common.HexToAddress("0x18ec44d1a96f99852a54b3a76628194c6169a795"))

	var address common.Address
	var tx *types.Transaction
	address, tx, evidenceFactory, err = factory_contract.DeployEvidenceFactory(fiscoClient.GetTransactOpts(), fiscoClient, addrs)
	zap.L().Info("deploy evidenceFactory contract success, ",
		zap.String("address", address.Hex()),
		zap.String("txid", tx.Hash().Hex()))
	return
}

// Close 手动关闭连接
func Close() {
	fiscoClient.Close()
}
