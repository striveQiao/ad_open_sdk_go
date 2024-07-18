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

// AdvertiserQualificationGetV30DataSubjectStatus
type AdvertiserQualificationGetV30DataSubjectStatus string

// List of advertiser_qualification_get_v3.0_data_subject_status
const (
	STATUS_CONFIRM_AdvertiserQualificationGetV30DataSubjectStatus         AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserQualificationGetV30DataSubjectStatus    AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserQualificationGetV30DataSubjectStatus      AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserQualificationGetV30DataSubjectStatus AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserQualificationGetV30DataSubjectStatus    AdvertiserQualificationGetV30DataSubjectStatus = "STATUS_WAIT_CONFIRM"
)

// All allowed values of AdvertiserQualificationGetV30DataSubjectStatus enum
var AllowedAdvertiserQualificationGetV30DataSubjectStatusEnumValues = []AdvertiserQualificationGetV30DataSubjectStatus{
	"STATUS_CONFIRM",
	"STATUS_CONFIRM_FAIL",
	"STATUS_NOT_SUBMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_CONFIRM",
}

// NewAdvertiserQualificationGetV30DataSubjectStatusFromValue returns a pointer to a valid AdvertiserQualificationGetV30DataSubjectStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationGetV30DataSubjectStatusFromValue(v string) (*AdvertiserQualificationGetV30DataSubjectStatus, error) {
	ev := AdvertiserQualificationGetV30DataSubjectStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationGetV30DataSubjectStatus: valid values are %v", v, AllowedAdvertiserQualificationGetV30DataSubjectStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationGetV30DataSubjectStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationGetV30DataSubjectStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_get_v3.0_data_subject_status value
func (v AdvertiserQualificationGetV30DataSubjectStatus) Ptr() *AdvertiserQualificationGetV30DataSubjectStatus {
	return &v
}
