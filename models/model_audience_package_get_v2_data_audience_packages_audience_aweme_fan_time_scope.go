/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope
type AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope string

// List of audience_package_get_v2_data_audience_packages_audience_aweme_fan_time_scope
const (
	FIFTEEN_DAYS_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope = "FIFTEEN_DAYS"
	SIXTY_DAYS_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope   AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope = "SIXTY_DAYS"
	THIRTY_DAYS_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope  AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope = "THIRTY_DAYS"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScopeEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope{
	"FIFTEEN_DAYS",
	"SIXTY_DAYS",
	"THIRTY_DAYS",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScopeFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScopeFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_aweme_fan_time_scope value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanTimeScope {
	return &v
}
