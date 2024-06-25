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

// AdvertiserQualificationCreateV2V2QualificationType
type AdvertiserQualificationCreateV2V2QualificationType string

// List of advertiser_qualification_create_v2_v2_qualification_type
const (
	QUALIFICATION_AD_AdvertiserQualificationCreateV2V2QualificationType AdvertiserQualificationCreateV2V2QualificationType = "QUALIFICATION_AD"
)

// All allowed values of AdvertiserQualificationCreateV2V2QualificationType enum
var AllowedAdvertiserQualificationCreateV2V2QualificationTypeEnumValues = []AdvertiserQualificationCreateV2V2QualificationType{
	"QUALIFICATION_AD",
}

// NewAdvertiserQualificationCreateV2V2QualificationTypeFromValue returns a pointer to a valid AdvertiserQualificationCreateV2V2QualificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationCreateV2V2QualificationTypeFromValue(v string) (*AdvertiserQualificationCreateV2V2QualificationType, error) {
	ev := AdvertiserQualificationCreateV2V2QualificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationCreateV2V2QualificationType: valid values are %v", v, AllowedAdvertiserQualificationCreateV2V2QualificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationCreateV2V2QualificationType) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationCreateV2V2QualificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_create_v2_v2_qualification_type value
func (v AdvertiserQualificationCreateV2V2QualificationType) Ptr() *AdvertiserQualificationCreateV2V2QualificationType {
	return &v
}
