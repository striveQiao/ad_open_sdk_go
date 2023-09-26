/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFileImageDeleteV10Request struct for QianchuanFileImageDeleteV10Request
type QianchuanFileImageDeleteV10Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 待删除的image_id列表，长度范围：1 ~ 100
	ImageIds []string `json:"image_ids"`
}
