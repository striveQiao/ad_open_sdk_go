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

// AdvertiserQualificationGetV30DataStatus
type AdvertiserQualificationGetV30DataStatus string

// List of advertiser_qualification_get_v3.0_data_status
const (
	STATUS_CONFIRM_AdvertiserQualificationGetV30DataStatus         AdvertiserQualificationGetV30DataStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserQualificationGetV30DataStatus    AdvertiserQualificationGetV30DataStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserQualificationGetV30DataStatus      AdvertiserQualificationGetV30DataStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserQualificationGetV30DataStatus AdvertiserQualificationGetV30DataStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserQualificationGetV30DataStatus    AdvertiserQualificationGetV30DataStatus = "STATUS_WAIT_CONFIRM"
)

// All allowed values of AdvertiserQualificationGetV30DataStatus enum
var AllowedAdvertiserQualificationGetV30DataStatusEnumValues = []AdvertiserQualificationGetV30DataStatus{
	"STATUS_CONFIRM",
	"STATUS_CONFIRM_FAIL",
	"STATUS_NOT_SUBMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_CONFIRM",
}

// NewAdvertiserQualificationGetV30DataStatusFromValue returns a pointer to a valid AdvertiserQualificationGetV30DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationGetV30DataStatusFromValue(v string) (*AdvertiserQualificationGetV30DataStatus, error) {
	ev := AdvertiserQualificationGetV30DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationGetV30DataStatus: valid values are %v", v, AllowedAdvertiserQualificationGetV30DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationGetV30DataStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationGetV30DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_get_v3.0_data_status value
func (v AdvertiserQualificationGetV30DataStatus) Ptr() *AdvertiserQualificationGetV30DataStatus {
	return &v
}
