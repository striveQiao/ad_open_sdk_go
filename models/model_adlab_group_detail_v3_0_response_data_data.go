/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataData 管家项目id
type AdlabGroupDetailV30ResponseDataData struct {
	// 计划id
	AdId   *int64                                     `json:"ad_id,omitempty"`
	AdInfo *AdlabGroupDetailV30ResponseDataDataAdInfo `json:"ad_info,omitempty"`
	// 广告组id
	CampaignId   int64                                           `json:"campaign_id"`
	CampaignInfo AdlabGroupDetailV30ResponseDataDataCampaignInfo `json:"campaign_info"`
	// 创意id
	CreativeId   *int64                                           `json:"creative_id,omitempty"`
	CreativeInfo *AdlabGroupDetailV30ResponseDataDataCreativeInfo `json:"creative_info,omitempty"`
	// 项目id
	Id int64 `json:"id"`
	// 项目名称
	Name      string                               `json:"name"`
	OptStatus AdlabGroupDetailV30DataDataOptStatus `json:"opt_status"`
	Status    AdlabGroupDetailV30DataDataStatus    `json:"status"`
}
