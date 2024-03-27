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

// AudiencePackageGetV2FilteringDeliveryRange
type AudiencePackageGetV2FilteringDeliveryRange string

// List of audience_package_get_v2_filtering_delivery_range
const (
	DEFAULT_AudiencePackageGetV2FilteringDeliveryRange   AudiencePackageGetV2FilteringDeliveryRange = "DEFAULT"
	UNION_AudiencePackageGetV2FilteringDeliveryRange     AudiencePackageGetV2FilteringDeliveryRange = "UNION"
	UNIVERSAL_AudiencePackageGetV2FilteringDeliveryRange AudiencePackageGetV2FilteringDeliveryRange = "UNIVERSAL"
)

// All allowed values of AudiencePackageGetV2FilteringDeliveryRange enum
var AllowedAudiencePackageGetV2FilteringDeliveryRangeEnumValues = []AudiencePackageGetV2FilteringDeliveryRange{
	"DEFAULT",
	"UNION",
	"UNIVERSAL",
}

// NewAudiencePackageGetV2FilteringDeliveryRangeFromValue returns a pointer to a valid AudiencePackageGetV2FilteringDeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2FilteringDeliveryRangeFromValue(v string) (*AudiencePackageGetV2FilteringDeliveryRange, error) {
	ev := AudiencePackageGetV2FilteringDeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2FilteringDeliveryRange: valid values are %v", v, AllowedAudiencePackageGetV2FilteringDeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2FilteringDeliveryRange) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2FilteringDeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_filtering_delivery_range value
func (v AudiencePackageGetV2FilteringDeliveryRange) Ptr() *AudiencePackageGetV2FilteringDeliveryRange {
	return &v
}
