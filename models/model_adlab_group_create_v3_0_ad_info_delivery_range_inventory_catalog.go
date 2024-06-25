/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog
type AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog string

// List of adlab_group_create_v3.0_ad_info_delivery_range_inventory_catalog
const (
	MANUAL_AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog          AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog = "MANUAL"
	UNIVERSAL_SMART_AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog = "UNIVERSAL_SMART"
)

// All allowed values of AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog enum
var AllowedAdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalogEnumValues = []AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog{
	"MANUAL",
	"UNIVERSAL_SMART",
}

// NewAdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalogFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalogFromValue(v string) (*AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog, error) {
	ev := AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalogEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalogEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_delivery_range_inventory_catalog value
func (v AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog) Ptr() *AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog {
	return &v
}
