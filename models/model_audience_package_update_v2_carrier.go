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

// AudiencePackageUpdateV2Carrier
type AudiencePackageUpdateV2Carrier string

// List of audience_package_update_v2_carrier
const (
	UNICOM_AudiencePackageUpdateV2Carrier AudiencePackageUpdateV2Carrier = "UNICOM"
	TELCOM_AudiencePackageUpdateV2Carrier AudiencePackageUpdateV2Carrier = "TELCOM"
	MOBILE_AudiencePackageUpdateV2Carrier AudiencePackageUpdateV2Carrier = "MOBILE"
)

// All allowed values of AudiencePackageUpdateV2Carrier enum
var AllowedAudiencePackageUpdateV2CarrierEnumValues = []AudiencePackageUpdateV2Carrier{
	"UNICOM",
	"TELCOM",
	"MOBILE",
}

// NewAudiencePackageUpdateV2CarrierFromValue returns a pointer to a valid AudiencePackageUpdateV2Carrier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2CarrierFromValue(v string) (*AudiencePackageUpdateV2Carrier, error) {
	ev := AudiencePackageUpdateV2Carrier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2Carrier: valid values are %v", v, AllowedAudiencePackageUpdateV2CarrierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2Carrier) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2CarrierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_carrier value
func (v AudiencePackageUpdateV2Carrier) Ptr() *AudiencePackageUpdateV2Carrier {
	return &v
}
