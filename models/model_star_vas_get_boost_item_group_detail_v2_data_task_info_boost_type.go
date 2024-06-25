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

// StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType
type StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType string

// List of star_vas_get_boost_item_group_detail_v2_data_task_info_boost_type
const (
	CUSTOM_PACKAGE_StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType = "CUSTOM_PACKAGE"
	CUSTOM_TAG_StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType     StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType = "CUSTOM_TAG"
	NONE_StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType           StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType = "NONE"
)

// All allowed values of StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType enum
var AllowedStarVasGetBoostItemGroupDetailV2DataTaskInfoBoostTypeEnumValues = []StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType{
	"CUSTOM_PACKAGE",
	"CUSTOM_TAG",
	"NONE",
}

// NewStarVasGetBoostItemGroupDetailV2DataTaskInfoBoostTypeFromValue returns a pointer to a valid StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarVasGetBoostItemGroupDetailV2DataTaskInfoBoostTypeFromValue(v string) (*StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType, error) {
	ev := StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType: valid values are %v", v, AllowedStarVasGetBoostItemGroupDetailV2DataTaskInfoBoostTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType) IsValid() bool {
	for _, existing := range AllowedStarVasGetBoostItemGroupDetailV2DataTaskInfoBoostTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_vas_get_boost_item_group_detail_v2_data_task_info_boost_type value
func (v StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType) Ptr() *StarVasGetBoostItemGroupDetailV2DataTaskInfoBoostType {
	return &v
}
