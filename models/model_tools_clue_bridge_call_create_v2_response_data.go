/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueBridgeCallCreateV2ResponseData
type ToolsClueBridgeCallCreateV2ResponseData struct {
	// 呼叫结果返回码
	CallResultCode *int64 `json:"call_result_code,omitempty"`
	// 呼叫结果返回信息
	CallResultMessage *string `json:"call_result_message,omitempty"`
	// 呼叫记录id, 用来标识一次外呼行为
	ContactId *string `json:"contact_id,omitempty"`
}
