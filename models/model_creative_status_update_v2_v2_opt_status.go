/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeStatusUpdateV2V2OptStatus
type CreativeStatusUpdateV2V2OptStatus string

// List of creative_status_update_v2_v2_opt_status
const (
	DELETE_CreativeStatusUpdateV2V2OptStatus  CreativeStatusUpdateV2V2OptStatus = "delete"
	ENABLE_CreativeStatusUpdateV2V2OptStatus  CreativeStatusUpdateV2V2OptStatus = "enable"
	DISABLE_CreativeStatusUpdateV2V2OptStatus CreativeStatusUpdateV2V2OptStatus = "disable"
)

// All allowed values of CreativeStatusUpdateV2V2OptStatus enum
var AllowedCreativeStatusUpdateV2V2OptStatusEnumValues = []CreativeStatusUpdateV2V2OptStatus{
	"delete",
	"enable",
	"disable",
}

// NewCreativeStatusUpdateV2V2OptStatusFromValue returns a pointer to a valid CreativeStatusUpdateV2V2OptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeStatusUpdateV2V2OptStatusFromValue(v string) (*CreativeStatusUpdateV2V2OptStatus, error) {
	ev := CreativeStatusUpdateV2V2OptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeStatusUpdateV2V2OptStatus: valid values are %v", v, AllowedCreativeStatusUpdateV2V2OptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeStatusUpdateV2V2OptStatus) IsValid() bool {
	for _, existing := range AllowedCreativeStatusUpdateV2V2OptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_status_update_v2_v2_opt_status value
func (v CreativeStatusUpdateV2V2OptStatus) Ptr() *CreativeStatusUpdateV2V2OptStatus {
	return &v
}
