/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType
type AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType string

// List of adlab_group_list_v3.0_data_groups_ad_info_track_url_setting_track_url_type
const (
	CUSTOM_AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType  AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType = "CUSTOM"
	EXISTED_AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType = "EXISTED"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType{
	"CUSTOM",
	"EXISTED",
}

// NewAdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_track_url_setting_track_url_type value
func (v AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType) Ptr() *AdlabGroupListV30DataGroupsAdInfoTrackUrlSettingTrackUrlType {
	return &v
}
