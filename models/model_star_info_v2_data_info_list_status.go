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

// StarInfoV2DataInfoListStatus
type StarInfoV2DataInfoListStatus string

// List of star_info_v2_data_info_list_status
const (
	DELETED_StarInfoV2DataInfoListStatus                    StarInfoV2DataInfoListStatus = "DELETED"
	ENABLE_StarInfoV2DataInfoListStatus                     StarInfoV2DataInfoListStatus = "ENABLE"
	FROZEN_StarInfoV2DataInfoListStatus                     StarInfoV2DataInfoListStatus = "FROZEN"
	NEW_PROTOCOL_StarInfoV2DataInfoListStatus               StarInfoV2DataInfoListStatus = "NEW_PROTOCOL"
	PUNISH_StarInfoV2DataInfoListStatus                     StarInfoV2DataInfoListStatus = "PUNISH"
	QUALIFICATION_VERIFICATION_StarInfoV2DataInfoListStatus StarInfoV2DataInfoListStatus = "QUALIFICATION_VERIFICATION"
	UN_PROTOCOL_StarInfoV2DataInfoListStatus                StarInfoV2DataInfoListStatus = "UN_PROTOCOL"
)

// All allowed values of StarInfoV2DataInfoListStatus enum
var AllowedStarInfoV2DataInfoListStatusEnumValues = []StarInfoV2DataInfoListStatus{
	"DELETED",
	"ENABLE",
	"FROZEN",
	"NEW_PROTOCOL",
	"PUNISH",
	"QUALIFICATION_VERIFICATION",
	"UN_PROTOCOL",
}

// NewStarInfoV2DataInfoListStatusFromValue returns a pointer to a valid StarInfoV2DataInfoListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarInfoV2DataInfoListStatusFromValue(v string) (*StarInfoV2DataInfoListStatus, error) {
	ev := StarInfoV2DataInfoListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarInfoV2DataInfoListStatus: valid values are %v", v, AllowedStarInfoV2DataInfoListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarInfoV2DataInfoListStatus) IsValid() bool {
	for _, existing := range AllowedStarInfoV2DataInfoListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_info_v2_data_info_list_status value
func (v StarInfoV2DataInfoListStatus) Ptr() *StarInfoV2DataInfoListStatus {
	return &v
}
