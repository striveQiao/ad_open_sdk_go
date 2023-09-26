/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserQualificationSubmitV30SubjectQualificationType
type AdvertiserQualificationSubmitV30SubjectQualificationType string

// List of advertiser_qualification_submit_v3.0_subject_qualification_type
const (
	ASSOCIATION_REGISTER_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType  AdvertiserQualificationSubmitV30SubjectQualificationType = "ASSOCIATION_REGISTER_CODE"
	COMMERCIAL_REGISTER_NUMBER_AdvertiserQualificationSubmitV30SubjectQualificationType AdvertiserQualificationSubmitV30SubjectQualificationType = "COMMERCIAL_REGISTER_NUMBER"
	COMPANY_CREDIT_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType        AdvertiserQualificationSubmitV30SubjectQualificationType = "COMPANY_CREDIT_CODE"
	COMPANY_REGISTER_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType      AdvertiserQualificationSubmitV30SubjectQualificationType = "COMPANY_REGISTER_CODE"
	CREDIT_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType                AdvertiserQualificationSubmitV30SubjectQualificationType = "CREDIT_CODE"
	HK_MACAO_TAIWAN_ID_AdvertiserQualificationSubmitV30SubjectQualificationType         AdvertiserQualificationSubmitV30SubjectQualificationType = "HK_MACAO_TAIWAN_ID"
	HK_REGISTER_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType           AdvertiserQualificationSubmitV30SubjectQualificationType = "HK_REGISTER_CODE"
	ID_AdvertiserQualificationSubmitV30SubjectQualificationType                         AdvertiserQualificationSubmitV30SubjectQualificationType = "ID"
	INDIVIDUAL_CREDIT_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType     AdvertiserQualificationSubmitV30SubjectQualificationType = "INDIVIDUAL_CREDIT_CODE"
	INDIVIDUAL_REGISTER_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType   AdvertiserQualificationSubmitV30SubjectQualificationType = "INDIVIDUAL_REGISTER_CODE"
	LAWYER_CERTIFICATE_NUMBER_AdvertiserQualificationSubmitV30SubjectQualificationType  AdvertiserQualificationSubmitV30SubjectQualificationType = "LAWYER_CERTIFICATE_NUMBER"
	LAWYER_PERMIT_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType         AdvertiserQualificationSubmitV30SubjectQualificationType = "LAWYER_PERMIT_CODE"
	LEGAL_PERSON_CREDIT_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType   AdvertiserQualificationSubmitV30SubjectQualificationType = "LEGAL_PERSON_CREDIT_CODE"
	ORGANIZATION_REGISTER_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType AdvertiserQualificationSubmitV30SubjectQualificationType = "ORGANIZATION_REGISTER_CODE"
	OTHER_AdvertiserQualificationSubmitV30SubjectQualificationType                      AdvertiserQualificationSubmitV30SubjectQualificationType = "OTHER"
	OVERSEA_REGISTER_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType      AdvertiserQualificationSubmitV30SubjectQualificationType = "OVERSEA_REGISTER_CODE"
	PASSPORT_ID_AdvertiserQualificationSubmitV30SubjectQualificationType                AdvertiserQualificationSubmitV30SubjectQualificationType = "PASSPORT_ID"
	SCHOOL_PERMIT_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType         AdvertiserQualificationSubmitV30SubjectQualificationType = "SCHOOL_PERMIT_CODE"
	SOCIAL_REGISTER_CODE_AdvertiserQualificationSubmitV30SubjectQualificationType       AdvertiserQualificationSubmitV30SubjectQualificationType = "SOCIAL_REGISTER_CODE"
)

// All allowed values of AdvertiserQualificationSubmitV30SubjectQualificationType enum
var AllowedAdvertiserQualificationSubmitV30SubjectQualificationTypeEnumValues = []AdvertiserQualificationSubmitV30SubjectQualificationType{
	"ASSOCIATION_REGISTER_CODE",
	"COMMERCIAL_REGISTER_NUMBER",
	"COMPANY_CREDIT_CODE",
	"COMPANY_REGISTER_CODE",
	"CREDIT_CODE",
	"HK_MACAO_TAIWAN_ID",
	"HK_REGISTER_CODE",
	"ID",
	"INDIVIDUAL_CREDIT_CODE",
	"INDIVIDUAL_REGISTER_CODE",
	"LAWYER_CERTIFICATE_NUMBER",
	"LAWYER_PERMIT_CODE",
	"LEGAL_PERSON_CREDIT_CODE",
	"ORGANIZATION_REGISTER_CODE",
	"OTHER",
	"OVERSEA_REGISTER_CODE",
	"PASSPORT_ID",
	"SCHOOL_PERMIT_CODE",
	"SOCIAL_REGISTER_CODE",
}

// NewAdvertiserQualificationSubmitV30SubjectQualificationTypeFromValue returns a pointer to a valid AdvertiserQualificationSubmitV30SubjectQualificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationSubmitV30SubjectQualificationTypeFromValue(v string) (*AdvertiserQualificationSubmitV30SubjectQualificationType, error) {
	ev := AdvertiserQualificationSubmitV30SubjectQualificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationSubmitV30SubjectQualificationType: valid values are %v", v, AllowedAdvertiserQualificationSubmitV30SubjectQualificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationSubmitV30SubjectQualificationType) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationSubmitV30SubjectQualificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_submit_v3.0_subject_qualification_type value
func (v AdvertiserQualificationSubmitV30SubjectQualificationType) Ptr() *AdvertiserQualificationSubmitV30SubjectQualificationType {
	return &v
}
