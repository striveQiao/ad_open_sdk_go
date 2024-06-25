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

// NativeAnchorGetV30FilteringStatus
type NativeAnchorGetV30FilteringStatus string

// List of native_anchor_get_v3.0_filtering_status
const (
	AUDIT_FAILD_NativeAnchorGetV30FilteringStatus   NativeAnchorGetV30FilteringStatus = "AUDIT_FAILD"
	AUDIT_SUCCESS_NativeAnchorGetV30FilteringStatus NativeAnchorGetV30FilteringStatus = "AUDIT_SUCCESS"
	CREATE_NativeAnchorGetV30FilteringStatus        NativeAnchorGetV30FilteringStatus = "CREATE"
	DELETE_NativeAnchorGetV30FilteringStatus        NativeAnchorGetV30FilteringStatus = "DELETE"
)

// All allowed values of NativeAnchorGetV30FilteringStatus enum
var AllowedNativeAnchorGetV30FilteringStatusEnumValues = []NativeAnchorGetV30FilteringStatus{
	"AUDIT_FAILD",
	"AUDIT_SUCCESS",
	"CREATE",
	"DELETE",
}

// NewNativeAnchorGetV30FilteringStatusFromValue returns a pointer to a valid NativeAnchorGetV30FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30FilteringStatusFromValue(v string) (*NativeAnchorGetV30FilteringStatus, error) {
	ev := NativeAnchorGetV30FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30FilteringStatus: valid values are %v", v, AllowedNativeAnchorGetV30FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30FilteringStatus) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_filtering_status value
func (v NativeAnchorGetV30FilteringStatus) Ptr() *NativeAnchorGetV30FilteringStatus {
	return &v
}
