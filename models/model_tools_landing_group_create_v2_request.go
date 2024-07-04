/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupCreateV2Request struct for ToolsLandingGroupCreateV2Request
type ToolsLandingGroupCreateV2Request struct {
	// 广告主id，范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserId       int64                                        `json:"advertiser_id"`
	ExperimentSiteType *ToolsLandingGroupCreateV2ExperimentSiteType `json:"experiment_site_type,omitempty"`
	GroupFlowType      ToolsLandingGroupCreateV2GroupFlowType       `json:"group_flow_type"`
	// 落地页组名称，范围：1 <= 长度 <= 20
	GroupTitle string `json:"group_title"`
	// 橙子建站站点id列表 ：2 <= 长度 <= 10 可通过【获取橙子建站站点列表】获取应答字段list
	SiteIds []string `json:"site_ids,omitempty"`
}
