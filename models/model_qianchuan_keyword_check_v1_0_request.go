/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanKeywordCheckV10Request struct for QianchuanKeywordCheckV10Request
type QianchuanKeywordCheckV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	// 关键词列表
	Keywords []string `json:"keywords"`
}
