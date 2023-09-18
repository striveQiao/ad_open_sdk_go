/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30ResponseDataGroupsInner struct for AdlabGroupListV30ResponseDataGroupsInner
type AdlabGroupListV30ResponseDataGroupsInner struct {
	// 计划id
	AdId   *int64                                          `json:"ad_id,omitempty"`
	AdInfo *AdlabGroupListV30ResponseDataGroupsInnerAdInfo `json:"ad_info,omitempty"`
	// 广告组id
	CampaignId   *int64                                                `json:"campaign_id,omitempty"`
	CampaignInfo *AdlabGroupListV30ResponseDataGroupsInnerCampaignInfo `json:"campaign_info,omitempty"`
	// 创意id
	CreativeId   *int64                                                `json:"creative_id,omitempty"`
	CreativeInfo *AdlabGroupListV30ResponseDataGroupsInnerCreativeInfo `json:"creative_info,omitempty"`
	// 管家项目id
	Id        *int64                                `json:"id,omitempty"`
	OptStatus *AdlabGroupListV30DataGroupsOptStatus `json:"opt_status,omitempty"`
	Status    *AdlabGroupListV30DataGroupsStatus    `json:"status,omitempty"`
}
