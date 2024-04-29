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

// AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes
type AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes string

// List of audience_package_get_v2_data_audience_packages_audience_action_v2_action_scenes
const (
	AD_AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes         AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes = "AD"
	APP_AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes        AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes = "APP"
	E_COMMERCE_AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes = "E-COMMERCE"
	NEWS_AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes       AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes = "NEWS"
	SEARCH_AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes     AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes = "SEARCH"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenesEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes{
	"AD",
	"APP",
	"E-COMMERCE",
	"NEWS",
	"SEARCH",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenesFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenesFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_action_v2_action_scenes value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceActionV2ActionScenes {
	return &v
}
