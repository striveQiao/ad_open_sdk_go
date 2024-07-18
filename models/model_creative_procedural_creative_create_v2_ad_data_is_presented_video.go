/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo
type CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo int64

// List of creative_procedural_creative_create_v2_ad_data_is_presented_video
const (
	Enum_0_CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo = 0
	Enum_1_CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo = 1
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo enum
var AllowedCreativeProceduralCreativeCreateV2AdDataIsPresentedVideoEnumValues = []CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo{
	0,
	1,
}

// NewCreativeProceduralCreativeCreateV2AdDataIsPresentedVideoFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataIsPresentedVideoFromValue(v int64) (*CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataIsPresentedVideoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataIsPresentedVideoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_is_presented_video value
func (v CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo) Ptr() *CreativeProceduralCreativeCreateV2AdDataIsPresentedVideo {
	return &v
}
