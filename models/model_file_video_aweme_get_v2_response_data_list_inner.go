/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAwemeGetV2ResponseDataListInner struct for FileVideoAwemeGetV2ResponseDataListInner
type FileVideoAwemeGetV2ResponseDataListInner struct {
	// 视频播放链接
	AwemePlayUrl *string `json:"aweme_play_url,omitempty"`
	// 视频时长，单位为秒
	Duration  *float64                              `json:"duration,omitempty"`
	ImageMode *FileVideoAwemeGetV2DataListImageMode `json:"image_mode,omitempty"`
	// 抖音视频ID
	ItemId *int64 `json:"item_id,omitempty"`
	// 素材ID
	Mid *int64 `json:"mid,omitempty"`
	// 视频标题
	Title *string `json:"title,omitempty"`
	// 视频封面ID
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	// 视频封面链接
	VideoCoverUrl *string `json:"video_cover_url,omitempty"`
	// 视频ID
	VideoId *string `json:"video_id,omitempty"`
}
