/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageUpdateV2LocationType
type AudiencePackageUpdateV2LocationType string

// List of audience_package_update_v2_location_type
const (
	CURRENT_AudiencePackageUpdateV2LocationType AudiencePackageUpdateV2LocationType = "CURRENT"
	HOME_AudiencePackageUpdateV2LocationType    AudiencePackageUpdateV2LocationType = "HOME"
	ALL_AudiencePackageUpdateV2LocationType     AudiencePackageUpdateV2LocationType = "ALL"
	TRAVEL_AudiencePackageUpdateV2LocationType  AudiencePackageUpdateV2LocationType = "TRAVEL"
)

// All allowed values of AudiencePackageUpdateV2LocationType enum
var AllowedAudiencePackageUpdateV2LocationTypeEnumValues = []AudiencePackageUpdateV2LocationType{
	"CURRENT",
	"HOME",
	"ALL",
	"TRAVEL",
}

// NewAudiencePackageUpdateV2LocationTypeFromValue returns a pointer to a valid AudiencePackageUpdateV2LocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2LocationTypeFromValue(v string) (*AudiencePackageUpdateV2LocationType, error) {
	ev := AudiencePackageUpdateV2LocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2LocationType: valid values are %v", v, AllowedAudiencePackageUpdateV2LocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2LocationType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2LocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_location_type value
func (v AudiencePackageUpdateV2LocationType) Ptr() *AudiencePackageUpdateV2LocationType {
	return &v
}
