/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus
type CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus int64

// List of creative_procedural_creative_create_v2_ad_data_ad_download_status
const (
	Enum_0_CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus = 0
	Enum_1_CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus = 1
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus enum
var AllowedCreativeProceduralCreativeCreateV2AdDataAdDownloadStatusEnumValues = []CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus{
	0,
	1,
}

// NewCreativeProceduralCreativeCreateV2AdDataAdDownloadStatusFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataAdDownloadStatusFromValue(v int64) (*CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataAdDownloadStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataAdDownloadStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_ad_download_status value
func (v CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus) Ptr() *CreativeProceduralCreativeCreateV2AdDataAdDownloadStatus {
	return &v
}
