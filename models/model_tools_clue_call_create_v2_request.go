/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallCreateV2Request struct for ToolsClueCallCreateV2Request
type ToolsClueCallCreateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 呼叫方式。1表示axb，3表示双呼，4表示网络电话。
	CallType int32 `json:"call_type"`
	// 线索号码
	CalleeNumber string `json:"callee_number"`
	// 客服电话。当呼叫方式为网络电话时，该字段可不传，其他类型均要传该字段。当呼叫方式为axb时，拨打中间号的手机号必须要与该号码一致。
	CallerNumber *string `json:"caller_number,omitempty"`
	// 线索ID
	ClueId int64 `json:"clue_id"`
	// 客服ID。平台下客服ID不得重复。
	UserId int64 `json:"user_id"`
}
