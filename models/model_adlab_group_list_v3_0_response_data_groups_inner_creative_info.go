/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30ResponseDataGroupsInnerCreativeInfo 创意维度信息
type AdlabGroupListV30ResponseDataGroupsInnerCreativeInfo struct {
	// 应用名
	AppName *string `json:"app_name,omitempty"`
	// 是否启用自动暂停规则
	AutoStop *int64 `json:"auto_stop,omitempty"`
	// 是否开启自动派生创意
	CreativeAutoGenerateSwitch *int64 `json:"creative_auto_generate_switch,omitempty"`
	// 落地页链接
	ExternalUrl *string `json:"external_url,omitempty"`
	// 创意分类-三级行业
	IndustryV3 *int64 `json:"industry_v3,omitempty"`
	// 是否开启评论
	IsCommentDisable *int64 `json:"is_comment_disable,omitempty"`
	// 是否跟随素材
	IsFollowMaterial *int64 `json:"is_follow_material,omitempty"`
	// 搭配试玩素材url
	PlayableUrl *string `json:"playable_url,omitempty"`
	// 广告来源
	Source *string `json:"source,omitempty"`
	// 应用下载详情页
	WebUrl *string `json:"web_url,omitempty"`
}
