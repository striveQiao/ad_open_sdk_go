/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeBannedCreateV30Request struct for ToolsAwemeBannedCreateV30Request
type ToolsAwemeBannedCreateV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 抖音号，当is_apply_to_adv不传或为false时，aweme_id生效
	AwemeId *string `json:"aweme_id,omitempty"`
	// 抖音用户id列表，传入屏蔽用户的抖音号； 当banned_type为AWEME_TYPE时必填；每次最多传50个抖音id，纯数字id不能以0开头；一个抖音号最多屏蔽5000个抖音ID。
	AwemeUserIds []string                             `json:"aweme_user_ids,omitempty"`
	BannedType   *ToolsAwemeBannedCreateV30BannedType `json:"banned_type,omitempty"`
	// 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
	IsApplyToAdv *bool `json:"is_apply_to_adv,omitempty"`
	// 抖音昵称关键词列表， 当banned_type为CUSTOM_TYPE时必填；关键词长度不大于20个字； 每次最多传50个关键词；
	NicknameKeywords []string `json:"nickname_keywords,omitempty"`
}
