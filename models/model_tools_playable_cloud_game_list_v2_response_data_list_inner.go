/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableCloudGameListV2ResponseDataListInner struct for ToolsPlayableCloudGameListV2ResponseDataListInner
type ToolsPlayableCloudGameListV2ResponseDataListInner struct {
	AdStatus *ToolsPlayableCloudGameListV2DataListAdStatus `json:"ad_status,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	EffectiveTimeEnd *string `json:"effective_time_end,omitempty"`
	//
	EffectiveTimeStart *string `json:"effective_time_start,omitempty"`
	//
	GameId *string `json:"game_id,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	Name        *string                                          `json:"name,omitempty"`
	Orientation *ToolsPlayableCloudGameListV2DataListOrientation `json:"orientation,omitempty"`
	//
	PreviewUrl *string                                     `json:"preview_url,omitempty"`
	Status     *ToolsPlayableCloudGameListV2DataListStatus `json:"status,omitempty"`
	//
	TrialTime *int64 `json:"trial_time,omitempty"`
}
