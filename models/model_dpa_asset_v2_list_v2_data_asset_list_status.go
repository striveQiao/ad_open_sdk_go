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

// DpaAssetV2ListV2DataAssetListStatus
type DpaAssetV2ListV2DataAssetListStatus string

// List of dpa_asset_v2_list_v2_data_asset_list_status
const (
	DISABLE_DpaAssetV2ListV2DataAssetListStatus DpaAssetV2ListV2DataAssetListStatus = "DISABLE"
	ENABLE_DpaAssetV2ListV2DataAssetListStatus  DpaAssetV2ListV2DataAssetListStatus = "ENABLE"
)

// All allowed values of DpaAssetV2ListV2DataAssetListStatus enum
var AllowedDpaAssetV2ListV2DataAssetListStatusEnumValues = []DpaAssetV2ListV2DataAssetListStatus{
	"DISABLE",
	"ENABLE",
}

// NewDpaAssetV2ListV2DataAssetListStatusFromValue returns a pointer to a valid DpaAssetV2ListV2DataAssetListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetV2ListV2DataAssetListStatusFromValue(v string) (*DpaAssetV2ListV2DataAssetListStatus, error) {
	ev := DpaAssetV2ListV2DataAssetListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetV2ListV2DataAssetListStatus: valid values are %v", v, AllowedDpaAssetV2ListV2DataAssetListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetV2ListV2DataAssetListStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetV2ListV2DataAssetListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_asset_v2_list_v2_data_asset_list_status value
func (v DpaAssetV2ListV2DataAssetListStatus) Ptr() *DpaAssetV2ListV2DataAssetListStatus {
	return &v
}
