/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// MaterialStatusUpdateV30Request struct for MaterialStatusUpdateV30Request
type MaterialStatusUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Data []*MaterialStatusUpdateV30RequestDataInner `json:"data"`
	//
	PromotionId int64 `json:"promotion_id"`
}
