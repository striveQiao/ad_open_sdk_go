/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoDeleteV2Request struct for FileVideoDeleteV2Request
type FileVideoDeleteV2Request struct {
	// 素材归属的广告主
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 待删除的素材ID列表，长度范围：1 ~ 100 注意：video_ids和material_ids至少填一个
	MaterialIds []int64 `json:"material_ids,omitempty"`
	// 待删除的视频ID列表，长度范围：1 ~ 100
	VideoIds []string `json:"video_ids,omitempty"`
}
