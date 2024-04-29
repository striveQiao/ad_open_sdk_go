/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType
type AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType string

// List of adlab_group_list_v3.0_data_groups_ad_info_delivery_range_union_video_type
const (
	GAME_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType = "GAME"
	NONE_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType = "NONE"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType{
	"GAME",
	"NONE",
}

// NewAdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_delivery_range_union_video_type value
func (v AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType) Ptr() *AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType {
	return &v
}
