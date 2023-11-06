/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRubeexGetV2Filtering
type ReportRubeexGetV2Filtering struct {
	//
	AdIds   []int64                            `json:"ad_ids,omitempty"`
	AppCode *ReportRubeexGetV2FilteringAppCode `json:"app_code,omitempty"`
	//
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
	//
	CreativeIds []int64 `json:"creative_ids,omitempty"`
	//
	DataMd5      *string                                 `json:"data_md5,omitempty"`
	DataScope    *ReportRubeexGetV2FilteringDataScope    `json:"data_scope,omitempty"`
	OsType       *ReportRubeexGetV2FilteringOsType       `json:"os_type,omitempty"`
	PlayableType *ReportRubeexGetV2FilteringPlayableType `json:"playable_type,omitempty"`
	Type         *ReportRubeexGetV2FilteringType         `json:"type,omitempty"`
}
