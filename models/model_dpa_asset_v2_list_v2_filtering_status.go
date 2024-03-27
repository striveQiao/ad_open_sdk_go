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

// DpaAssetV2ListV2FilteringStatus
type DpaAssetV2ListV2FilteringStatus string

// List of dpa_asset_v2_list_v2_filtering_status
const (
	DISABLE_DpaAssetV2ListV2FilteringStatus DpaAssetV2ListV2FilteringStatus = "DISABLE"
	ENABLE_DpaAssetV2ListV2FilteringStatus  DpaAssetV2ListV2FilteringStatus = "ENABLE"
)

// All allowed values of DpaAssetV2ListV2FilteringStatus enum
var AllowedDpaAssetV2ListV2FilteringStatusEnumValues = []DpaAssetV2ListV2FilteringStatus{
	"DISABLE",
	"ENABLE",
}

// NewDpaAssetV2ListV2FilteringStatusFromValue returns a pointer to a valid DpaAssetV2ListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetV2ListV2FilteringStatusFromValue(v string) (*DpaAssetV2ListV2FilteringStatus, error) {
	ev := DpaAssetV2ListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetV2ListV2FilteringStatus: valid values are %v", v, AllowedDpaAssetV2ListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetV2ListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetV2ListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_asset_v2_list_v2_filtering_status value
func (v DpaAssetV2ListV2FilteringStatus) Ptr() *DpaAssetV2ListV2FilteringStatus {
	return &v
}
