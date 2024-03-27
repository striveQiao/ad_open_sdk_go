/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderReplyAuthorCancelV2Request struct for StarOrderReplyAuthorCancelV2Request
type StarOrderReplyAuthorCancelV2Request struct {
	// (1)同意；(2)拒绝
	Accept int64 `json:"accept"`
	// 任务ID
	OrderId int64 `json:"order_id"`
	// 拒绝理由
	RejectReason *string `json:"reject_reason,omitempty"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}
