/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthListV2ResponseDataListInnerVideoInfo
type ToolsAwemeAuthListV2ResponseDataListInnerVideoInfo struct {
	// 视频播放链接
	AwemePlayUrl *string `json:"aweme_play_url,omitempty"`
	// 视频时长，单位为秒
	Duration  *float64                                        `json:"duration,omitempty"`
	ImageMode *ToolsAwemeAuthListV2DataListVideoInfoImageMode `json:"image_mode,omitempty"`
	// 抖音视频ID
	ItemId int64 `json:"item_id"`
	// 素材ID
	Mid *int64 `json:"mid,omitempty"`
	// 视频名称
	Title *string `json:"title,omitempty"`
	// 视频封面ID
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	// 视频封面链接
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`
	// 视频ID
	VideoId *string `json:"video_id,omitempty"`
}
