/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeCreateV2Request struct for CreativeCustomCreativeCreateV2Request
type CreativeCustomCreativeCreateV2Request struct {
	AdData CreativeCustomCreativeCreateV2RequestAdData `json:"ad_data"`
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CreativeList []*CreativeCustomCreativeCreateV2RequestCreativeListInner `json:"creative_list"`
}
