package router

import (
	"fisco-cert-app/controller"
	"github.com/gin-gonic/gin"
)

func BlockChainRouter(r *gin.RouterGroup) {

	r.GET("blockNumber", controller.BlockHeightHandler)
	r.GET("getPeers", controller.BlockGetPeersHandler)
	r.GET("getTransactionTotal", controller.TxTotalHandler)
	r.GET("getLatestTx", controller.TxLatestDetailHandler)
	r.GET("getTransSums", controller.TxTenDaySumsHandler)
	r.GET("getLatestBlock", controller.BlockLatestHandler)
}
