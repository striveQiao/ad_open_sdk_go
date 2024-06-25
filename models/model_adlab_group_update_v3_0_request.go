/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateV30Request struct for AdlabGroupUpdateV30Request
type AdlabGroupUpdateV30Request struct {
	AdInfo *AdlabGroupUpdateV30RequestAdInfo `json:"ad_info,omitempty"`
	// 广告主id
	AdvertiserId int64                                   `json:"advertiser_id"`
	CampaignInfo *AdlabGroupUpdateV30RequestCampaignInfo `json:"campaign_info,omitempty"`
	CreativeInfo *AdlabGroupUpdateV30RequestCreativeInfo `json:"creative_info,omitempty"`
	// 管家项目id
	Id int64 `json:"id"`
}
