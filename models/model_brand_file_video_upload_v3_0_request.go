/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandFileVideoUploadV30Request struct for BrandFileVideoUploadV30Request
type BrandFileVideoUploadV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 视频文件
	VideoFile *FormFileInfo `json:"video_file"`
}
