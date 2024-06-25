/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableListGetV2ResponseDataListInner struct for ToolsPlayableListGetV2ResponseDataListInner
type ToolsPlayableListGetV2ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	PlayableId *int64 `json:"playable_id,omitempty"`
	//
	PlayableName        *string                                            `json:"playable_name,omitempty"`
	PlayableOrientation *ToolsPlayableListGetV2DataListPlayableOrientation `json:"playable_orientation,omitempty"`
	//
	PlayableUrl *string                               `json:"playable_url,omitempty"`
	Status      *ToolsPlayableListGetV2DataListStatus `json:"status,omitempty"`
}
