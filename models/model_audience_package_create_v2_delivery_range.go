/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2DeliveryRange
type AudiencePackageCreateV2DeliveryRange string

// List of audience_package_create_v2_delivery_range
const (
	UNION_AudiencePackageCreateV2DeliveryRange     AudiencePackageCreateV2DeliveryRange = "UNION"
	DEFAULT_AudiencePackageCreateV2DeliveryRange   AudiencePackageCreateV2DeliveryRange = "DEFAULT"
	UNIVERSAL_AudiencePackageCreateV2DeliveryRange AudiencePackageCreateV2DeliveryRange = "UNIVERSAL"
)

// All allowed values of AudiencePackageCreateV2DeliveryRange enum
var AllowedAudiencePackageCreateV2DeliveryRangeEnumValues = []AudiencePackageCreateV2DeliveryRange{
	"UNION",
	"DEFAULT",
	"UNIVERSAL",
}

// NewAudiencePackageCreateV2DeliveryRangeFromValue returns a pointer to a valid AudiencePackageCreateV2DeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2DeliveryRangeFromValue(v string) (*AudiencePackageCreateV2DeliveryRange, error) {
	ev := AudiencePackageCreateV2DeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2DeliveryRange: valid values are %v", v, AllowedAudiencePackageCreateV2DeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2DeliveryRange) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2DeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_delivery_range value
func (v AudiencePackageCreateV2DeliveryRange) Ptr() *AudiencePackageCreateV2DeliveryRange {
	return &v
}
