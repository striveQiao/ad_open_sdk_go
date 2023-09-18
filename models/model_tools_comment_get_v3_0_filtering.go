/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentGetV30Filtering
type ToolsCommentGetV30Filtering struct {
	// 计划id列表，一次最多100个，仅当platform_version=V1时传入生效
	AdIds        []int64                                  `json:"ad_ids,omitempty"`
	BindRelation *ToolsCommentGetV30FilteringBindRelation `json:"bind_relation,omitempty"`
	// 广告创意id列表，一次最多100个，仅当platform_version=V1时传入生效
	CreativeIds []int64                                 `json:"creative_ids,omitempty"`
	EmotionType *ToolsCommentGetV30FilteringEmotionType `json:"emotion_type,omitempty"`
	HideStatus  *ToolsCommentGetV30FilteringHideStatus  `json:"hide_status,omitempty"`
	// 是否已回复
	IsReplied *int64 `json:"is_replied,omitempty"`
	// 广告视频id列表，一次最多100个，仅当platform_version=V2时传入生效
	ItemIds   []int64                               `json:"item_ids,omitempty"`
	LevelType *ToolsCommentGetV30FilteringLevelType `json:"level_type,omitempty"`
}
