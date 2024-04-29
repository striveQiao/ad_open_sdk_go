/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordCampaignUpdateV2Request struct for ToolsPrivativeWordCampaignUpdateV2Request
type ToolsPrivativeWordCampaignUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CampaignId int64 `json:"campaign_id"`
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
}
