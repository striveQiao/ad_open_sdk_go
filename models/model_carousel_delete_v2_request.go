/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselDeleteV2Request struct for CarouselDeleteV2Request
type CarouselDeleteV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	// 图集 id
	CarouselIds []int64 `json:"carousel_ids"`
}
