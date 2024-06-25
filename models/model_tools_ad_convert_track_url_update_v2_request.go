/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdConvertTrackUrlUpdateV2Request struct for ToolsAdConvertTrackUrlUpdateV2Request
type ToolsAdConvertTrackUrlUpdateV2Request struct {
	//
	ActionTrackUrl *string `json:"action_track_url,omitempty"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ConvertId *int64 `json:"convert_id,omitempty"`
	//
	TrackUrl *string `json:"track_url,omitempty"`
	//
	VideoPlayDoneTrackUrl *string `json:"video_play_done_track_url,omitempty"`
	//
	VideoPlayEffectiveTrackUrl *string `json:"video_play_effective_track_url,omitempty"`
	//
	VideoPlayTrackUrl *string `json:"video_play_track_url,omitempty"`
}
