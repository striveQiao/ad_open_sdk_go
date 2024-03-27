/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRtaGetV2ResponseDataListInner struct for ReportRtaGetV2ResponseDataListInner
type ReportRtaGetV2ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CreativeId *int64 `json:"creative_id,omitempty"`
	//
	IsRtaBid *int64                                      `json:"is_rta_bid,omitempty"`
	Metrics  *ReportRtaGetV2ResponseDataListInnerMetrics `json:"metrics,omitempty"`
	//
	Platform *string `json:"platform,omitempty"`
	//
	StatTimeDay *string `json:"stat_time_day,omitempty"`
}
