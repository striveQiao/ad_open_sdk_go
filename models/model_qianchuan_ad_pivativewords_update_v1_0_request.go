/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdPivativewordsUpdateV10Request struct for QianchuanAdPivativewordsUpdateV10Request
type QianchuanAdPivativewordsUpdateV10Request struct {
	// 计划ID
	AdId int64 `json:"ad_id"`
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 短语否词，传空数组为清空
	PhraseWords []string `json:"phrase_words,omitempty"`
	// 精确否词，传空数组为清空
	PreciseWords []string `json:"precise_words,omitempty"`
}
