/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV30Request struct for KeywordCreateV30Request
type KeywordCreateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Keywords []*KeywordCreateV30RequestKeywordsInner `json:"keywords"`
	//
	PromotionId int64 `json:"promotion_id"`
}
