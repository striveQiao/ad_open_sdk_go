/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallCreateV2ResponseData
type ToolsClueCallCreateV2ResponseData struct {
	// 呼叫结果返回码。[详见“呼叫结果返回码 & 呼叫结果返回信息”](https://bytedance.feishu.cn/wiki/wikcneKh4vQDeYrIN33joWdlZCd#)
	CallResultCode *int32 `json:"call_result_code,omitempty"`
	// 呼叫结果返回信息。[详见“呼叫结果返回码 & 呼叫结果返回信息”](https://bytedance.feishu.cn/wiki/wikcneKh4vQDeYrIN33joWdlZCd#)
	CallResultMessage *string `json:"call_result_message,omitempty"`
	// 呼叫记录id, 用来标识一次外呼行为
	ContactId *string `json:"contact_id,omitempty"`
	// 中间号。当呼叫方式为axb时该字段有值。
	VirtualNumber *string `json:"virtual_number,omitempty"`
}
