package models

import "time"

type AuthResp struct {
	Mark    string `json:"watermark_url"`
	Address string `json:"address"`
}

type EvidenceDetailResp struct {
	//*EvidenceResp
	*EvidenceResp
	Evidence  string    `json:"evidence"`
	CreatedAt time.Time `json:"create_at"`
	// 增加简介（该图片介绍）
}

type EvidenceResp struct {
	TxID      string    `json:"tx_id"`      // 版权编号/交易id
	ImgURL    string    `json:"img_url"`    // 图片
	UserName  string    `json:"user_name"`  // 用户名
	CreatedAt time.Time `json:"created_at"` // 时间
	Total     int64     `json:"total"`      // 总数量
}

type TxTotalResp struct {
	TxSum       string `json:"tx_sum"`        // 交易总数
	FailedTxSum string `json:"failed_tx_sum"` // 失败交易总数
}

type TxDetailResp struct {
	From      string `json:"from"`
	To        string `json:"to"`
	TxID      string `json:"tx_id"`
	Timestamp string `json:"timestamp"`
}

type BlockDetailResp struct {
	Height    string `json:"height"`    // 当前块高
	Sealer    string `json:"sealer"`    // 出块节点
	TxSum     string `json:"tx_sum"`    // 交易总数
	Timestamp string `json:"timestamp"` // 时间戳
}
