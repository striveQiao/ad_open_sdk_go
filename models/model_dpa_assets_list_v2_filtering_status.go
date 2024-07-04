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

// DpaAssetsListV2FilteringStatus
type DpaAssetsListV2FilteringStatus int64

// List of dpa_assets_list_v2_filtering_status
const (
	Enum_0_DpaAssetsListV2FilteringStatus DpaAssetsListV2FilteringStatus = 0
	Enum_1_DpaAssetsListV2FilteringStatus DpaAssetsListV2FilteringStatus = 1
)

// All allowed values of DpaAssetsListV2FilteringStatus enum
var AllowedDpaAssetsListV2FilteringStatusEnumValues = []DpaAssetsListV2FilteringStatus{
	0,
	1,
}

// NewDpaAssetsListV2FilteringStatusFromValue returns a pointer to a valid DpaAssetsListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsListV2FilteringStatusFromValue(v int64) (*DpaAssetsListV2FilteringStatus, error) {
	ev := DpaAssetsListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsListV2FilteringStatus: valid values are %v", v, AllowedDpaAssetsListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetsListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_list_v2_filtering_status value
func (v DpaAssetsListV2FilteringStatus) Ptr() *DpaAssetsListV2FilteringStatus {
	return &v
}
