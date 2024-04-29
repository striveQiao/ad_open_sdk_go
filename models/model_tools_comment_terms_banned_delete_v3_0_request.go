/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentTermsBannedDeleteV30Request struct for ToolsCommentTermsBannedDeleteV30Request
type ToolsCommentTermsBannedDeleteV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 抖音号，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeId *string `json:"aweme_id,omitempty"`
	// 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv *bool `json:"is_apply_to_adv,omitempty"`
	// 屏蔽词列表，一次最多操作100个词，屏蔽词最大20字
	Terms []string `json:"terms"`
}
