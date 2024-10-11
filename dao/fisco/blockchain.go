package fisco

import (
	"context"
	"fisco-cert-app/models"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func GetBlockHeight() int64 {
	height, _ := fiscoClient.GetBlockNumber(context.Background())
	return height
}

func GetNodes() ([]string, error) {
	nodes, err := fiscoClient.GetNodeIDList(context.Background())

	if err != nil {
		zap.L().Error("fiscoClient.GetPeers failed", zap.Error(err))
		return nil, err
	}
	str := string(nodes)
	slices := strings.Split(str, ",")
	// 将每个子字符串转换为对应的字节切片
	result := make([]string, len(slices))
	for i, s := range slices {
		for _, char := range s {
			if (int(char) >= int('a') && int(char) <= int('z')) || (int(char) >= int('0') && int(char) <= int('9')) {
				result[i] += string(char)
			}
		}
	}
	return result, nil
}

func TxTotal() (total *models.TxTotalResp, err error) {
	txCount, err := fiscoClient.GetTotalTransactionCount(context.Background())
	if err != nil {
		zap.L().Error("fiscoClient.GetTotalTransactionCount failed", zap.Error(err))
		return nil, err
	}
	txSum, _ := strconv.ParseInt(strings.Replace(txCount.GetTxSum(), "0x", "", 1), 16, 64)
	failTxSum, _ := strconv.ParseInt(strings.Replace(txCount.GetFailedTxSum(), "0x", "", 1), 16, 64)
	return &models.TxTotalResp{
		TxSum:       fmt.Sprint(txSum),
		FailedTxSum: fmt.Sprint(failTxSum),
	}, nil
}

// 获取最新十个区块
func GetBlockByPage() ([]*models.BlockDetailResp, error) {
	var blocks = make([]*models.BlockDetailResp, 0)

	blockHeight := GetBlockHeight()
	ctx := context.Background()
	nodeIds, err := GetNodes()
	if err != nil {
		zap.L().Error("fiscoClient getBlockByNumber failed", zap.Error(err))
		return nil, err
	}
	if blockHeight < 10 {
		for i := blockHeight; i >= 0; i-- {
			nowBlock, errByFiscoNumber := fiscoClient.GetBlockByNumber(ctx, i, false)
			if errByFiscoNumber != nil {
				zap.L().Error("fiscoClient getBlockByNumber failed", zap.Error(errByFiscoNumber))
				return nil, errByFiscoNumber
			}
			// 拿到节点索引
			nodeIndex, _ := strconv.ParseInt(strings.Replace(nowBlock.GetSealer(), "0x", "", 1), 16, 64)
			// 倒序索引
			revertIndex := int64(len(nodeIds)) - nodeIndex - 1
			block := &models.BlockDetailResp{
				Height:    fmt.Sprint(i),
				Sealer:    nodeIds[revertIndex],
				TxSum:     fmt.Sprint(len(nowBlock.Transactions)),
				Timestamp: nowBlock.Timestamp,
			}

			blocks = append(blocks, block)
		}
	} else {
		for i := blockHeight; i >= blockHeight-10; i-- {
			nowBlock, errByFiscoNumber := fiscoClient.GetBlockByNumber(ctx, i, false)
			if errByFiscoNumber != nil {
				zap.L().Error("fiscoClient getBlockByNumber failed", zap.Error(errByFiscoNumber))
				return nil, errByFiscoNumber
			}
			// 拿到节点索引
			nodeIndex, _ := strconv.ParseInt(strings.Replace(nowBlock.GetSealer(), "0x", "", 1), 16, 64)
			// 倒序索引
			revertIndex := int64(len(nodeIds)) - nodeIndex - 1
			block := &models.BlockDetailResp{
				Height:    fmt.Sprint(i),
				Sealer:    nodeIds[revertIndex],
				TxSum:     fmt.Sprint(len(nowBlock.Transactions)),
				Timestamp: nowBlock.Timestamp,
			}

			blocks = append(blocks, block)
		}
	}

	return blocks, nil
}

func GetTransactionByPage() ([]*models.TxDetailResp, error) {
	var txs = make([]*models.TxDetailResp, 0)
	blockHeight := GetBlockHeight()
	ctx := context.Background()
	total, _ := TxTotal()
	txSum, _ := strconv.ParseInt(total.TxSum, 16, 64)
	if txSum < 10 { // 小于10到创世块即停止
		for i := blockHeight; i >= 0; i-- {
			block, err := fiscoClient.GetBlockByNumber(ctx, i, false)
			if err != nil {
				zap.L().Info("fiscoClient.GetBlockByNumber failed", zap.Error(err))
				continue
			}
			nowTxSum := len(block.GetTransactions())
			for idx := 0; idx < nowTxSum; idx++ {
				txDetail, errByFisco := fiscoClient.GetTransactionByBlockNumberAndIndex(ctx, i, idx)
				if errByFisco != nil {
					zap.L().Info("fiscoClient.GetTransactionByBlockNumberAndIndex failed", zap.Error(err))
					continue
				}
				timestamp, _ := strconv.ParseInt(strings.Replace(block.Timestamp, "0x", "", 1), 16, 64)
				txDetailResp := &models.TxDetailResp{
					From:      txDetail.From,
					To:        txDetail.To,
					TxID:      txDetail.Hash,
					Timestamp: fmt.Sprint(timestamp),
				}
				txs = append(txs, txDetailResp)
			}
		}
	} else {
		// 计算十笔交易，由于每个块的交易数量都不一致，需要从后往前不断计算
		sum := 0
		for i := blockHeight; i >= 0; i-- {
			block, err := fiscoClient.GetBlockByNumber(ctx, i, false)
			if err != nil {
				zap.L().Info("fiscoClient.GetBlockByNumber failed", zap.Error(err))
				continue
			}
			nowTxSum := len(block.GetTransactions())
			for idx := 0; idx < nowTxSum; idx++ {
				sum += 1
				txDetail, errByFisco := fiscoClient.GetTransactionByBlockNumberAndIndex(ctx, i, idx)
				if errByFisco != nil {
					zap.L().Info("fiscoClient.GetTransactionByBlockNumberAndIndex failed", zap.Error(err))
					continue
				}
				timestamp, _ := strconv.ParseInt(strings.Replace(block.Timestamp, "0x", "", 1), 16, 64)
				txDetailResp := &models.TxDetailResp{
					From:      txDetail.From,
					To:        txDetail.To,
					TxID:      txDetail.Hash,
					Timestamp: fmt.Sprint(timestamp),
				}
				txs = append(txs, txDetailResp)

				if sum >= 10 {
					return txs, errByFisco
				}
			}
		}
	}
	return txs, nil
}
