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

// AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType
type AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType string

// List of advertiser_delivery_pkg_config_v3.0_data_industry_config_unnecessaries_rules_type
const (
	CHOICE_AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType  AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType = "CHOICE"
	COMPOSE_AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType = "COMPOSE"
)

// All allowed values of AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType enum
var AllowedAdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesTypeEnumValues = []AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType{
	"CHOICE",
	"COMPOSE",
}

// NewAdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesTypeFromValue returns a pointer to a valid AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesTypeFromValue(v string) (*AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType, error) {
	ev := AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType: valid values are %v", v, AllowedAdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType) IsValid() bool {
	for _, existing := range AllowedAdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_delivery_pkg_config_v3.0_data_industry_config_unnecessaries_rules_type value
func (v AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType) Ptr() *AdvertiserDeliveryPkgConfigV30DataIndustryConfigUnnecessariesRulesType {
	return &v
}
