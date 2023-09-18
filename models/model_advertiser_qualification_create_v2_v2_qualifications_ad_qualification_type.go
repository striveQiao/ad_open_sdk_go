/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserQualificationCreateV2V2QualificationsAdQualificationType
type AdvertiserQualificationCreateV2V2QualificationsAdQualificationType string

// List of advertiser_qualification_create_v2_v2_qualifications_ad_qualification_type
const (
	AUTHORIZATION_AdvertiserQualificationCreateV2V2QualificationsAdQualificationType  AdvertiserQualificationCreateV2V2QualificationsAdQualificationType = "AUTHORIZATION"
	CERTIFY_AdvertiserQualificationCreateV2V2QualificationsAdQualificationType        AdvertiserQualificationCreateV2V2QualificationsAdQualificationType = "CERTIFY"
	TRADEMARK_CERT_AdvertiserQualificationCreateV2V2QualificationsAdQualificationType AdvertiserQualificationCreateV2V2QualificationsAdQualificationType = "TRADEMARK_CERT"
	AD_DATA_CERT_AdvertiserQualificationCreateV2V2QualificationsAdQualificationType   AdvertiserQualificationCreateV2V2QualificationsAdQualificationType = "AD_DATA_CERT"
	INSPECT_REPORT_AdvertiserQualificationCreateV2V2QualificationsAdQualificationType AdvertiserQualificationCreateV2V2QualificationsAdQualificationType = "INSPECT_REPORT"
)

// All allowed values of AdvertiserQualificationCreateV2V2QualificationsAdQualificationType enum
var AllowedAdvertiserQualificationCreateV2V2QualificationsAdQualificationTypeEnumValues = []AdvertiserQualificationCreateV2V2QualificationsAdQualificationType{
	"AUTHORIZATION",
	"CERTIFY",
	"TRADEMARK_CERT",
	"AD_DATA_CERT",
	"INSPECT_REPORT",
}

// NewAdvertiserQualificationCreateV2V2QualificationsAdQualificationTypeFromValue returns a pointer to a valid AdvertiserQualificationCreateV2V2QualificationsAdQualificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationCreateV2V2QualificationsAdQualificationTypeFromValue(v string) (*AdvertiserQualificationCreateV2V2QualificationsAdQualificationType, error) {
	ev := AdvertiserQualificationCreateV2V2QualificationsAdQualificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationCreateV2V2QualificationsAdQualificationType: valid values are %v", v, AllowedAdvertiserQualificationCreateV2V2QualificationsAdQualificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationCreateV2V2QualificationsAdQualificationType) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationCreateV2V2QualificationsAdQualificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_create_v2_v2_qualifications_ad_qualification_type value
func (v AdvertiserQualificationCreateV2V2QualificationsAdQualificationType) Ptr() *AdvertiserQualificationCreateV2V2QualificationsAdQualificationType {
	return &v
}
