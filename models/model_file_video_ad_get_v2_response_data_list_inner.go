/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAdGetV2ResponseDataListInner struct for FileVideoAdGetV2ResponseDataListInner
type FileVideoAdGetV2ResponseDataListInner struct {
	// 码率，单位bps
	BitRate *float64 `json:"bit_rate,omitempty"`
	// 视频时长
	Duration *float64 `json:"duration,omitempty"`
	// 视频格式
	Format *string `json:"format,omitempty"`
	// 视频高度
	Height *int64 `json:"height,omitempty"`
	// 视频ID
	Id *string `json:"id,omitempty"`
	// 素材id，即多合一报表中的素材id，一个素材唯一对应一个素材id
	MaterialId *int64 `json:"material_id,omitempty"`
	// 视频首帧截图
	PosterUrl *string `json:"poster_url,omitempty"`
	// 视频md5值
	Signature *string `json:"signature,omitempty"`
	// 视频大小
	Size *int64 `json:"size,omitempty"`
	// 视频地址
	Url *string `json:"url,omitempty"`
	// 视频宽度
	Width *int64 `json:"width,omitempty"`
}
