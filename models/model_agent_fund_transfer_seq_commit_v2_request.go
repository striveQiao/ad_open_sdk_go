/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentFundTransferSeqCommitV2Request struct for AgentFundTransferSeqCommitV2Request
type AgentFundTransferSeqCommitV2Request struct {
	// 代理商ID
	AgentId int64 `json:"agent_id"`
	// 转账流水号，第一步请求会返回
	TransferSeq int64 `json:"transfer_seq"`
}
