/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2Request struct for KeywordUpdateV2V2Request
type KeywordUpdateV2V2Request struct {
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Keywords []*KeywordUpdateV2V2RequestKeywordsInner `json:"keywords"`
}
