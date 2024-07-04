/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatDataGetV2ResponseDataData
type ClueWechatDataGetV2ResponseDataData struct {
	// 计划id
	AdId *int64 `json:"ad_id,omitempty"`
	// 渠道来源
	Channel *string `json:"channel,omitempty"`
	//
	ClickId *string `json:"click_id,omitempty"`
	// 转化时间，格式\"YYYY-MM-DD hh:mm:ss\"
	ConvertTime *string `json:"convert_time,omitempty"`
	// 创意id
	CreativeId *int64 `json:"creative_id,omitempty"`
	// 销售人员id
	FollowUserIds []string `json:"follow_user_ids,omitempty"`
	// 加粉用户好友昵称
	WechatUserName *string `json:"wechat_user_name,omitempty"`
}
