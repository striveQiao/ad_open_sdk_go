/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus
type SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus string

// List of shared_wallet_wallet_info_get_v3.0_data_wallet_info_common_wallet_info_shared_wallet_status
const (
	DEFUALT_SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus = "DEFUALT"
)

// All allowed values of SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus enum
var AllowedSharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatusEnumValues = []SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus{
	"DEFUALT",
}

// NewSharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatusFromValue returns a pointer to a valid SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatusFromValue(v string) (*SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus, error) {
	ev := SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus: valid values are %v", v, AllowedSharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus) IsValid() bool {
	for _, existing := range AllowedSharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_wallet_info_get_v3.0_data_wallet_info_common_wallet_info_shared_wallet_status value
func (v SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus) Ptr() *SharedWalletWalletInfoGetV30DataWalletInfoCommonWalletInfoSharedWalletStatus {
	return &v
}
