/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeStatusUpdateV10Request struct for QianchuanCreativeStatusUpdateV10Request
type QianchuanCreativeStatusUpdateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CreativeIds []int64                                   `json:"creative_ids"`
	OptStatus   QianchuanCreativeStatusUpdateV10OptStatus `json:"opt_status"`
}
