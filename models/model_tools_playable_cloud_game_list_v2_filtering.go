/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableCloudGameListV2Filtering
type ToolsPlayableCloudGameListV2Filtering struct {
	//
	AdStatus []*ToolsPlayableCloudGameListV2FilteringAdStatus `json:"ad_status,omitempty"`
	//
	GameIds []string `json:"game_ids,omitempty"`
	//
	Ids []int64 `json:"ids,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Orientation []*ToolsPlayableCloudGameListV2FilteringOrientation `json:"orientation,omitempty"`
	//
	Status []*ToolsPlayableCloudGameListV2FilteringStatus `json:"status,omitempty"`
}
