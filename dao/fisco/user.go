package fisco

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// GenUser 生成非国密用户
func GenUser() (address, hexPrivateKey string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	hexPrivateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", "", errors.New("publicKey is not of type *ecdsa.PublicKey")
	}

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return
}

// equalsHexPrivateKeyAddress 私钥比对
func EqualsHexPrivateKeyAddress(hexPrivateKey, addr string) (flag bool, err error) {
	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		return false, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return false, errors.New("publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return addr == address, nil
}
