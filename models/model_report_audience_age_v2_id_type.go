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

// ReportAudienceAgeV2IdType
type ReportAudienceAgeV2IdType string

// List of report_audience_age_v2_id_type
const (
	AUDIENCE_STAT_ID_TYPE_AD_ReportAudienceAgeV2IdType         ReportAudienceAgeV2IdType = "AUDIENCE_STAT_ID_TYPE_AD"
	AUDIENCE_STAT_ID_TYPE_ADVERTISER_ReportAudienceAgeV2IdType ReportAudienceAgeV2IdType = "AUDIENCE_STAT_ID_TYPE_ADVERTISER"
	AUDIENCE_STAT_ID_TYPE_CAMPAIGN_ReportAudienceAgeV2IdType   ReportAudienceAgeV2IdType = "AUDIENCE_STAT_ID_TYPE_CAMPAIGN"
)

// All allowed values of ReportAudienceAgeV2IdType enum
var AllowedReportAudienceAgeV2IdTypeEnumValues = []ReportAudienceAgeV2IdType{
	"AUDIENCE_STAT_ID_TYPE_AD",
	"AUDIENCE_STAT_ID_TYPE_ADVERTISER",
	"AUDIENCE_STAT_ID_TYPE_CAMPAIGN",
}

// NewReportAudienceAgeV2IdTypeFromValue returns a pointer to a valid ReportAudienceAgeV2IdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceAgeV2IdTypeFromValue(v string) (*ReportAudienceAgeV2IdType, error) {
	ev := ReportAudienceAgeV2IdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceAgeV2IdType: valid values are %v", v, AllowedReportAudienceAgeV2IdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceAgeV2IdType) IsValid() bool {
	for _, existing := range AllowedReportAudienceAgeV2IdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_age_v2_id_type value
func (v ReportAudienceAgeV2IdType) Ptr() *ReportAudienceAgeV2IdType {
	return &v
}
