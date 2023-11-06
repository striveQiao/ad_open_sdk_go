/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserQualificationSubmitV30SubjectCheckType
type AdvertiserQualificationSubmitV30SubjectCheckType string

// List of advertiser_qualification_submit_v3.0_subject_check_type
const (
	COMPANY_AdvertiserQualificationSubmitV30SubjectCheckType         AdvertiserQualificationSubmitV30SubjectCheckType = "COMPANY"
	GOVERNMENT_AdvertiserQualificationSubmitV30SubjectCheckType      AdvertiserQualificationSubmitV30SubjectCheckType = "GOVERNMENT"
	HK_MACAO_TAIWAN_AdvertiserQualificationSubmitV30SubjectCheckType AdvertiserQualificationSubmitV30SubjectCheckType = "HK_MACAO_TAIWAN"
	INDIVIDUAL_AdvertiserQualificationSubmitV30SubjectCheckType      AdvertiserQualificationSubmitV30SubjectCheckType = "INDIVIDUAL"
	OTHERS_AdvertiserQualificationSubmitV30SubjectCheckType          AdvertiserQualificationSubmitV30SubjectCheckType = "OTHERS"
	OVERSEA_AdvertiserQualificationSubmitV30SubjectCheckType         AdvertiserQualificationSubmitV30SubjectCheckType = "OVERSEA"
	SELF_EMPLOY_AdvertiserQualificationSubmitV30SubjectCheckType     AdvertiserQualificationSubmitV30SubjectCheckType = "SELF_EMPLOY"
)

// All allowed values of AdvertiserQualificationSubmitV30SubjectCheckType enum
var AllowedAdvertiserQualificationSubmitV30SubjectCheckTypeEnumValues = []AdvertiserQualificationSubmitV30SubjectCheckType{
	"COMPANY",
	"GOVERNMENT",
	"HK_MACAO_TAIWAN",
	"INDIVIDUAL",
	"OTHERS",
	"OVERSEA",
	"SELF_EMPLOY",
}

// NewAdvertiserQualificationSubmitV30SubjectCheckTypeFromValue returns a pointer to a valid AdvertiserQualificationSubmitV30SubjectCheckType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationSubmitV30SubjectCheckTypeFromValue(v string) (*AdvertiserQualificationSubmitV30SubjectCheckType, error) {
	ev := AdvertiserQualificationSubmitV30SubjectCheckType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationSubmitV30SubjectCheckType: valid values are %v", v, AllowedAdvertiserQualificationSubmitV30SubjectCheckTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationSubmitV30SubjectCheckType) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationSubmitV30SubjectCheckTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_submit_v3.0_subject_check_type value
func (v AdvertiserQualificationSubmitV30SubjectCheckType) Ptr() *AdvertiserQualificationSubmitV30SubjectCheckType {
	return &v
}
