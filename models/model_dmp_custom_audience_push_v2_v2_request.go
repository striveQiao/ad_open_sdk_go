/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpCustomAudiencePushV2V2Request struct for DmpCustomAudiencePushV2V2Request
type DmpCustomAudiencePushV2V2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CustomAudienceId int64 `json:"custom_audience_id"`
	//
	TargetAdvertiserIds []int64 `json:"target_advertiser_ids"`
}
