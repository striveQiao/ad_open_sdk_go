/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentTransferTransactionRecordV2ResponseData
type AgentTransferTransactionRecordV2ResponseData struct {
	PageInfo *AgentTransferTransactionRecordV2ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	Records []*AgentTransferTransactionRecordV2ResponseDataRecordsInner `json:"records"`
}
