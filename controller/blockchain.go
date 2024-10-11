package controller

import (
	"fisco-cert-app/dao/fisco"
	"fisco-cert-app/dao/mysql"
	"fisco-cert-app/logic"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// BlockHeightHandler 获取最新块高
// @Summary 获取最新块高
// @Description 获取fisco-bcos最新块高
// @Tags 区块链信息相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string "height"
// @Router /blockNumber [get]
func BlockHeightHandler(c *gin.Context) {
	ResponseSuccess(c, logic.GetBlockHeight())
}

// BlockGetPeersHandler 最新区块链连接的节点ID
// @Summary 最新区块链连接的节点ID
// @Description 获取fisco-bcos区块链连接的节点ID
// @Tags 区块链信息相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []string "data"
// @Router /getPeers [get]
func BlockGetPeersHandler(c *gin.Context) {
	nodeIds, err := fisco.GetNodes()
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err)
		return
	}
	ResponseSuccess(c, nodeIds)
}

// TxTotalHandler 区块链交易总数
// @Summary 最新区块链交易总数
// @Description 获取fisco-bcos区块链交易总数
// @Tags 区块链信息相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} models.TxTotalResp "data"
// @Router /getTransactionTotal [get]
func TxTotalHandler(c *gin.Context) {
	txResp, err := fisco.TxTotal()
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err)
		return
	}
	ResponseSuccess(c, txResp)
}

// BlockLatestHandler 获取最新十个区块
// @Summary 获取最新十个区块
// @Description 获取fisco-bcos区块链获取最新十个区块
// @Tags 区块链信息相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []models.BlockDetailResp "data"
// @Router /getLatestBlock [get]
func BlockLatestHandler(c *gin.Context) {
	blocks, err := fisco.GetBlockByPage()
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err)
		return
	}
	ResponseSuccess(c, blocks)
}

// TxLatestDetailHandler 最新是十笔交易
// @Summary 获取最新十个十笔交易
// @Description 获取fisco-bcos区块链获取最新十笔区块
// @Tags 区块链信息相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []models.TxDetailResp "data"
// @Router /getLatestTx [get]
func TxLatestDetailHandler(c *gin.Context) {
	txs, err := fisco.GetTransactionByPage()
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err)
		return
	}
	ResponseSuccess(c, txs)
}

// 查询近十日的交易数量
// @Summary 获取通过水印系统的交易数量
// @Description 获取印系统的交易数量最新用户交易数量
// @Tags 区块链信息相关接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []mysql.TxSum "data"
// @Router /getTransSums [get]
func TxTenDaySumsHandler(c *gin.Context) {
	sums, err := mysql.GetNowTenTxCount()
	fmt.Println(sums)
	if err != nil {
		zap.L().Error("mysql.GetNowTenTxCount failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err)
		return
	}
	ResponseSuccess(c, sums)
}
