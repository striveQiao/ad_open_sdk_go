/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2FilteringPlatformVersion
type ReportAdvertiserGetV2FilteringPlatformVersion string

// List of report_advertiser_get_v2_filtering_platform_version
const (
	V1_ReportAdvertiserGetV2FilteringPlatformVersion  ReportAdvertiserGetV2FilteringPlatformVersion = "V1"
	V2_ReportAdvertiserGetV2FilteringPlatformVersion  ReportAdvertiserGetV2FilteringPlatformVersion = "V2"
	ALL_ReportAdvertiserGetV2FilteringPlatformVersion ReportAdvertiserGetV2FilteringPlatformVersion = "ALL"
)

// All allowed values of ReportAdvertiserGetV2FilteringPlatformVersion enum
var AllowedReportAdvertiserGetV2FilteringPlatformVersionEnumValues = []ReportAdvertiserGetV2FilteringPlatformVersion{
	"V1",
	"V2",
	"ALL",
}

// NewReportAdvertiserGetV2FilteringPlatformVersionFromValue returns a pointer to a valid ReportAdvertiserGetV2FilteringPlatformVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2FilteringPlatformVersionFromValue(v string) (*ReportAdvertiserGetV2FilteringPlatformVersion, error) {
	ev := ReportAdvertiserGetV2FilteringPlatformVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2FilteringPlatformVersion: valid values are %v", v, AllowedReportAdvertiserGetV2FilteringPlatformVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2FilteringPlatformVersion) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2FilteringPlatformVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_filtering_platform_version value
func (v ReportAdvertiserGetV2FilteringPlatformVersion) Ptr() *ReportAdvertiserGetV2FilteringPlatformVersion {
	return &v
}
