/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignUpdateStatusV2Request struct for CampaignUpdateStatusV2Request
type CampaignUpdateStatusV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CampaignIds []int64                         `json:"campaign_ids"`
	OptStatus   CampaignUpdateStatusV2OptStatus `json:"opt_status"`
}
