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

// CreativeCustomCreativeUpdateV2AdDataPriorityTrial
type CreativeCustomCreativeUpdateV2AdDataPriorityTrial string

// List of creative_custom_creative_update_v2_ad_data_priority_trial
const (
	ON_CreativeCustomCreativeUpdateV2AdDataPriorityTrial  CreativeCustomCreativeUpdateV2AdDataPriorityTrial = "ON"
	OFF_CreativeCustomCreativeUpdateV2AdDataPriorityTrial CreativeCustomCreativeUpdateV2AdDataPriorityTrial = "OFF"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataPriorityTrial enum
var AllowedCreativeCustomCreativeUpdateV2AdDataPriorityTrialEnumValues = []CreativeCustomCreativeUpdateV2AdDataPriorityTrial{
	"ON",
	"OFF",
}

// NewCreativeCustomCreativeUpdateV2AdDataPriorityTrialFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataPriorityTrial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataPriorityTrialFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataPriorityTrial, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataPriorityTrial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataPriorityTrial: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataPriorityTrialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataPriorityTrial) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataPriorityTrialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_priority_trial value
func (v CreativeCustomCreativeUpdateV2AdDataPriorityTrial) Ptr() *CreativeCustomCreativeUpdateV2AdDataPriorityTrial {
	return &v
}
