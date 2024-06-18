/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesLandingType
type AudiencePackageGetV2DataAudiencePackagesLandingType string

// List of audience_package_get_v2_data_audience_packages_landing_type
const (
	APP_AudiencePackageGetV2DataAudiencePackagesLandingType         AudiencePackageGetV2DataAudiencePackagesLandingType = "APP"
	APP_ANDROID_AudiencePackageGetV2DataAudiencePackagesLandingType AudiencePackageGetV2DataAudiencePackagesLandingType = "APP_ANDROID"
	APP_IOS_AudiencePackageGetV2DataAudiencePackagesLandingType     AudiencePackageGetV2DataAudiencePackagesLandingType = "APP_IOS"
	ARTICLE_AudiencePackageGetV2DataAudiencePackagesLandingType     AudiencePackageGetV2DataAudiencePackagesLandingType = "ARTICLE"
	AWEME_AudiencePackageGetV2DataAudiencePackagesLandingType       AudiencePackageGetV2DataAudiencePackagesLandingType = "AWEME"
	DPA_AudiencePackageGetV2DataAudiencePackagesLandingType         AudiencePackageGetV2DataAudiencePackagesLandingType = "DPA"
	EXTERNAL_AudiencePackageGetV2DataAudiencePackagesLandingType    AudiencePackageGetV2DataAudiencePackagesLandingType = "EXTERNAL"
	GOODS_AudiencePackageGetV2DataAudiencePackagesLandingType       AudiencePackageGetV2DataAudiencePackagesLandingType = "GOODS"
	LIVE_AudiencePackageGetV2DataAudiencePackagesLandingType        AudiencePackageGetV2DataAudiencePackagesLandingType = "LIVE"
	MICRO_GAME_AudiencePackageGetV2DataAudiencePackagesLandingType  AudiencePackageGetV2DataAudiencePackagesLandingType = "MICRO_GAME"
	QUICK_APP_AudiencePackageGetV2DataAudiencePackagesLandingType   AudiencePackageGetV2DataAudiencePackagesLandingType = "QUICK_APP"
	SHOP_AudiencePackageGetV2DataAudiencePackagesLandingType        AudiencePackageGetV2DataAudiencePackagesLandingType = "SHOP"
	STORE_AudiencePackageGetV2DataAudiencePackagesLandingType       AudiencePackageGetV2DataAudiencePackagesLandingType = "STORE"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesLandingType enum
var AllowedAudiencePackageGetV2DataAudiencePackagesLandingTypeEnumValues = []AudiencePackageGetV2DataAudiencePackagesLandingType{
	"APP",
	"APP_ANDROID",
	"APP_IOS",
	"ARTICLE",
	"AWEME",
	"DPA",
	"EXTERNAL",
	"GOODS",
	"LIVE",
	"MICRO_GAME",
	"QUICK_APP",
	"SHOP",
	"STORE",
}

// NewAudiencePackageGetV2DataAudiencePackagesLandingTypeFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesLandingTypeFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesLandingType, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesLandingType: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesLandingType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_landing_type value
func (v AudiencePackageGetV2DataAudiencePackagesLandingType) Ptr() *AudiencePackageGetV2DataAudiencePackagesLandingType {
	return &v
}
