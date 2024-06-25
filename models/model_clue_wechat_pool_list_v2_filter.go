/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatPoolListV2Filter
type ClueWechatPoolListV2Filter struct {
	// 审核状态，0审核中、1审核拒绝、2审核通过
	AuditStatus []*ClueWechatPoolListV2FilterAuditStatus `json:"audit_status,omitempty"`
	// 微信号码包ID
	InstanceId *int64 `json:"instance_id,omitempty"`
	// 微信昵称，模糊匹配
	NickName *string `json:"nick_name,omitempty"`
	// 分页，>=1，默认值1
	Page *int32 `json:"page,omitempty"`
	// 页大小，1-100，默认值20
	PageSize *int32 `json:"page_size,omitempty"`
	// 微信号，模糊匹配
	WechatNumber *string `json:"wechat_number,omitempty"`
}
