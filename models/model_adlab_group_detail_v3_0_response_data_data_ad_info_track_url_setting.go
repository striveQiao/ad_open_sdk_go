/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataAdInfoTrackUrlSetting 监测链接设置
type AdlabGroupDetailV30ResponseDataDataAdInfoTrackUrlSetting struct {
	// 点击（监测链接）
	ActionTrackUrl []string `json:"action_track_url,omitempty"`
	// 展示（监测链接）
	TrackUrl []string `json:"track_url,omitempty"`
	// 监测链接组id
	TrackUrlGroupId  *int64                                                            `json:"track_url_group_id,omitempty"`
	TrackUrlSendType *AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlSendType `json:"track_url_send_type,omitempty"`
	TrackUrlType     *AdlabGroupDetailV30DataDataAdInfoTrackUrlSettingTrackUrlType     `json:"track_url_type,omitempty"`
	// 视频播完（监测链接）
	VideoPlayDoneTrackUrl []string `json:"video_play_done_track_url,omitempty"`
	// 视频有效播放（监测链接）
	VideoPlayEffectiveTrackUrl []string `json:"video_play_effective_track_url,omitempty"`
	// 视频开始播放（监测链接）
	VideoPlayTrackUrl []string `json:"video_play_track_url,omitempty"`
}
