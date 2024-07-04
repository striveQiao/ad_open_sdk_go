/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgDeleteV30Request struct for AdvertiserDeliveryPkgDeleteV30Request
type AdvertiserDeliveryPkgDeleteV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 推广产品pkg_id列表，支持传入多个批量删除，list长度1~50
	PkgIds []int64 `json:"pkg_ids"`
}
