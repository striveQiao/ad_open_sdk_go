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

// DpaAssetV2DetailReadV2DataAssetListStatus
type DpaAssetV2DetailReadV2DataAssetListStatus string

// List of dpa_asset_v2_detail_read_v2_data_asset_list_status
const (
	DISABLE_DpaAssetV2DetailReadV2DataAssetListStatus DpaAssetV2DetailReadV2DataAssetListStatus = "DISABLE"
	ENABLE_DpaAssetV2DetailReadV2DataAssetListStatus  DpaAssetV2DetailReadV2DataAssetListStatus = "ENABLE"
)

// All allowed values of DpaAssetV2DetailReadV2DataAssetListStatus enum
var AllowedDpaAssetV2DetailReadV2DataAssetListStatusEnumValues = []DpaAssetV2DetailReadV2DataAssetListStatus{
	"DISABLE",
	"ENABLE",
}

// NewDpaAssetV2DetailReadV2DataAssetListStatusFromValue returns a pointer to a valid DpaAssetV2DetailReadV2DataAssetListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetV2DetailReadV2DataAssetListStatusFromValue(v string) (*DpaAssetV2DetailReadV2DataAssetListStatus, error) {
	ev := DpaAssetV2DetailReadV2DataAssetListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetV2DetailReadV2DataAssetListStatus: valid values are %v", v, AllowedDpaAssetV2DetailReadV2DataAssetListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetV2DetailReadV2DataAssetListStatus) IsValid() bool {
	for _, existing := range AllowedDpaAssetV2DetailReadV2DataAssetListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_asset_v2_detail_read_v2_data_asset_list_status value
func (v DpaAssetV2DetailReadV2DataAssetListStatus) Ptr() *DpaAssetV2DetailReadV2DataAssetListStatus {
	return &v
}
