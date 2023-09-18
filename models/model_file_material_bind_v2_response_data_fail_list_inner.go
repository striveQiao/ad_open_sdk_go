/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialBindV2ResponseDataFailListInner struct for FileMaterialBindV2ResponseDataFailListInner
type FileMaterialBindV2ResponseDataFailListInner struct {
	// 推送失败的图集id
	CarouselId *int64 `json:"carousel_id,omitempty"`
	// 推送失败原因
	FailReason *string `json:"fail_reason,omitempty"`
	// 推送失败的图片id
	ImageId *string `json:"image_id,omitempty"`
	// 目标推送广告主id
	TargetAdvertiserId *int64 `json:"target_advertiser_id,omitempty"`
	// 推送失败的视频id
	VideoId *string `json:"video_id,omitempty"`
}
