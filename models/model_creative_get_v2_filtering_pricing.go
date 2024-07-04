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

// CreativeGetV2FilteringPricing
type CreativeGetV2FilteringPricing string

// List of creative_get_v2_filtering_pricing
const (
	PRICING_CPA_CreativeGetV2FilteringPricing      CreativeGetV2FilteringPricing = "PRICING_CPA"
	PRICING_CPC_CreativeGetV2FilteringPricing      CreativeGetV2FilteringPricing = "PRICING_CPC"
	PRICING_CPC_OCPM_CreativeGetV2FilteringPricing CreativeGetV2FilteringPricing = "PRICING_CPC_OCPM"
	PRICING_CPM_CreativeGetV2FilteringPricing      CreativeGetV2FilteringPricing = "PRICING_CPM"
	PRICING_CPV_CreativeGetV2FilteringPricing      CreativeGetV2FilteringPricing = "PRICING_CPV"
	PRICING_OCPC_CreativeGetV2FilteringPricing     CreativeGetV2FilteringPricing = "PRICING_OCPC"
	PRICING_OCPM_CreativeGetV2FilteringPricing     CreativeGetV2FilteringPricing = "PRICING_OCPM"
)

// All allowed values of CreativeGetV2FilteringPricing enum
var AllowedCreativeGetV2FilteringPricingEnumValues = []CreativeGetV2FilteringPricing{
	"PRICING_CPA",
	"PRICING_CPC",
	"PRICING_CPC_OCPM",
	"PRICING_CPM",
	"PRICING_CPV",
	"PRICING_OCPC",
	"PRICING_OCPM",
}

// NewCreativeGetV2FilteringPricingFromValue returns a pointer to a valid CreativeGetV2FilteringPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeGetV2FilteringPricingFromValue(v string) (*CreativeGetV2FilteringPricing, error) {
	ev := CreativeGetV2FilteringPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeGetV2FilteringPricing: valid values are %v", v, AllowedCreativeGetV2FilteringPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeGetV2FilteringPricing) IsValid() bool {
	for _, existing := range AllowedCreativeGetV2FilteringPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_get_v2_filtering_pricing value
func (v CreativeGetV2FilteringPricing) Ptr() *CreativeGetV2FilteringPricing {
	return &v
}
