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

// AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus
type AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus string

// List of advertiser_delivery_pkg_config_v3.0_data_industry_config_industry_status
const (
	NONVALID_AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus = "NONVALID"
	VALID_AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus    AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus = "VALID"
)

// All allowed values of AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus enum
var AllowedAdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatusEnumValues = []AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus{
	"NONVALID",
	"VALID",
}

// NewAdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatusFromValue returns a pointer to a valid AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatusFromValue(v string) (*AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus, error) {
	ev := AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus: valid values are %v", v, AllowedAdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_delivery_pkg_config_v3.0_data_industry_config_industry_status value
func (v AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus) Ptr() *AdvertiserDeliveryPkgConfigV30DataIndustryConfigIndustryStatus {
	return &v
}
