/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdUpdateStatusV2Request struct for AdUpdateStatusV2Request
type AdUpdateStatusV2Request struct {
	//
	AdIds []int64 `json:"ad_ids"`
	//
	AdvertiserId int64                     `json:"advertiser_id"`
	OptStatus    AdUpdateStatusV2OptStatus `json:"opt_status"`
}
