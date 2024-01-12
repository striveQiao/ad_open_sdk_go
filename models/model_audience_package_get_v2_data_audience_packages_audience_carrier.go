/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesAudienceCarrier
type AudiencePackageGetV2DataAudiencePackagesAudienceCarrier string

// List of audience_package_get_v2_data_audience_packages_audience_carrier
const (
	MOBILE_AudiencePackageGetV2DataAudiencePackagesAudienceCarrier AudiencePackageGetV2DataAudiencePackagesAudienceCarrier = "MOBILE"
	TELCOM_AudiencePackageGetV2DataAudiencePackagesAudienceCarrier AudiencePackageGetV2DataAudiencePackagesAudienceCarrier = "TELCOM"
	UNICOM_AudiencePackageGetV2DataAudiencePackagesAudienceCarrier AudiencePackageGetV2DataAudiencePackagesAudienceCarrier = "UNICOM"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceCarrier enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceCarrierEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceCarrier{
	"MOBILE",
	"TELCOM",
	"UNICOM",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceCarrierFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceCarrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceCarrierFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceCarrier, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceCarrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceCarrier: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceCarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceCarrier) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceCarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_carrier value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceCarrier) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceCarrier {
	return &v
}
