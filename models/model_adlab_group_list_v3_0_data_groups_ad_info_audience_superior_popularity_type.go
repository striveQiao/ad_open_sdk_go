/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType
type AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType string

// List of adlab_group_list_v3.0_data_groups_ad_info_audience_superior_popularity_type
const (
	GAME_AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType = "GAME"
	NONE_AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType = "NONE"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewAdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_audience_superior_popularity_type value
func (v AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType) Ptr() *AdlabGroupListV30DataGroupsAdInfoAudienceSuperiorPopularityType {
	return &v
}
