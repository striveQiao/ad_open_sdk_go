/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatPoolListV2ResponseDataItemsInner struct for ClueWechatPoolListV2ResponseDataItemsInner
type ClueWechatPoolListV2ResponseDataItemsInner struct {
	// 公众号的开发者ID
	AppId       *string                                   `json:"app_id,omitempty"`
	AuditStatus *ClueWechatPoolListV2DataItemsAuditStatus `json:"audit_status,omitempty"`
	// 复制次数
	ClueCount *int64 `json:"clue_count,omitempty"`
	// 创建时间，格式：2022-07-21 15:36:09
	CreateTime *string                                 `json:"create_time,omitempty"`
	HasQrCode  *ClueWechatPoolListV2DataItemsHasQrCode `json:"has_qr_code,omitempty"`
	// 生效号码包数量
	InstanceCount *int64 `json:"instance_count,omitempty"`
	// 修改时间，格式：2022-07-21 15:36:09
	ModTime *string `json:"mod_time,omitempty"`
	// 微信昵称
	NickName *string `json:"nick_name,omitempty"`
	// 审核拒绝理由
	RejectReason *string `json:"reject_reason,omitempty"`
	// 启用状态  0启用 1禁用
	Status *int64 `json:"status,omitempty"`
	// 尾缀
	Suffix *string `json:"suffix,omitempty"`
	// 微信号，企业微信时返回为空
	WechatNumber *string                                  `json:"wechat_number,omitempty"`
	WechatType   *ClueWechatPoolListV2DataItemsWechatType `json:"wechat_type,omitempty"`
}
