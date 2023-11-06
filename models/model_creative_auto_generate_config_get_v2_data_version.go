/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeAutoGenerateConfigGetV2DataVersion
type CreativeAutoGenerateConfigGetV2DataVersion string

// List of creative_auto_generate_config_get_v2_data_version
const (
	STRATEGY_CreativeAutoGenerateConfigGetV2DataVersion CreativeAutoGenerateConfigGetV2DataVersion = "Strategy"
	TEMPLATE_CreativeAutoGenerateConfigGetV2DataVersion CreativeAutoGenerateConfigGetV2DataVersion = "Template"
)

// All allowed values of CreativeAutoGenerateConfigGetV2DataVersion enum
var AllowedCreativeAutoGenerateConfigGetV2DataVersionEnumValues = []CreativeAutoGenerateConfigGetV2DataVersion{
	"Strategy",
	"Template",
}

// NewCreativeAutoGenerateConfigGetV2DataVersionFromValue returns a pointer to a valid CreativeAutoGenerateConfigGetV2DataVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeAutoGenerateConfigGetV2DataVersionFromValue(v string) (*CreativeAutoGenerateConfigGetV2DataVersion, error) {
	ev := CreativeAutoGenerateConfigGetV2DataVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeAutoGenerateConfigGetV2DataVersion: valid values are %v", v, AllowedCreativeAutoGenerateConfigGetV2DataVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeAutoGenerateConfigGetV2DataVersion) IsValid() bool {
	for _, existing := range AllowedCreativeAutoGenerateConfigGetV2DataVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_auto_generate_config_get_v2_data_version value
func (v CreativeAutoGenerateConfigGetV2DataVersion) Ptr() *CreativeAutoGenerateConfigGetV2DataVersion {
	return &v
}
