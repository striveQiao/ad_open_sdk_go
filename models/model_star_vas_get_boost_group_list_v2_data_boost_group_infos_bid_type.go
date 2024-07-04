/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarVasGetBoostGroupListV2DataBoostGroupInfosBidType
type StarVasGetBoostGroupListV2DataBoostGroupInfosBidType string

// List of star_vas_get_boost_group_list_v2_data_boost_group_infos_bid_type
const (
	AUTO_StarVasGetBoostGroupListV2DataBoostGroupInfosBidType   StarVasGetBoostGroupListV2DataBoostGroupInfosBidType = "AUTO"
	MANUAL_StarVasGetBoostGroupListV2DataBoostGroupInfosBidType StarVasGetBoostGroupListV2DataBoostGroupInfosBidType = "MANUAL"
)

// All allowed values of StarVasGetBoostGroupListV2DataBoostGroupInfosBidType enum
var AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosBidTypeEnumValues = []StarVasGetBoostGroupListV2DataBoostGroupInfosBidType{
	"AUTO",
	"MANUAL",
}

// NewStarVasGetBoostGroupListV2DataBoostGroupInfosBidTypeFromValue returns a pointer to a valid StarVasGetBoostGroupListV2DataBoostGroupInfosBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarVasGetBoostGroupListV2DataBoostGroupInfosBidTypeFromValue(v string) (*StarVasGetBoostGroupListV2DataBoostGroupInfosBidType, error) {
	ev := StarVasGetBoostGroupListV2DataBoostGroupInfosBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarVasGetBoostGroupListV2DataBoostGroupInfosBidType: valid values are %v", v, AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosBidType) IsValid() bool {
	for _, existing := range AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_vas_get_boost_group_list_v2_data_boost_group_infos_bid_type value
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosBidType) Ptr() *StarVasGetBoostGroupListV2DataBoostGroupInfosBidType {
	return &v
}
