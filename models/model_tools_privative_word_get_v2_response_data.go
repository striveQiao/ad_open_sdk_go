/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordGetV2ResponseData
type ToolsPrivativeWordGetV2ResponseData struct {
	//
	AdErrorList []int64 `json:"ad_error_list,omitempty"`
	//
	AdsPrivative []*ToolsPrivativeWordGetV2ResponseDataAdsPrivativeInner `json:"ads_privative,omitempty"`
	//
	CampaignErrorList []int64 `json:"campaign_error_list,omitempty"`
	//
	CampaignsPrivative []*ToolsPrivativeWordGetV2ResponseDataCampaignsPrivativeInner `json:"campaigns_privative,omitempty"`
}
