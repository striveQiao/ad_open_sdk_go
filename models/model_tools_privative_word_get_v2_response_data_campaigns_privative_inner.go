/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordGetV2ResponseDataCampaignsPrivativeInner struct for ToolsPrivativeWordGetV2ResponseDataCampaignsPrivativeInner
type ToolsPrivativeWordGetV2ResponseDataCampaignsPrivativeInner struct {
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	PhraseWords []string `json:"phrase_words,omitempty"`
	//
	PreciseWords []string `json:"precise_words,omitempty"`
}
