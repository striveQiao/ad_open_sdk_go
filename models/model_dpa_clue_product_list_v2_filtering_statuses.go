/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaClueProductListV2FilteringStatuses
type DpaClueProductListV2FilteringStatuses string

// List of dpa_clue_product_list_v2_filtering_statuses
const (
	STATUS_OFFLINE_DpaClueProductListV2FilteringStatuses DpaClueProductListV2FilteringStatuses = "STATUS_OFFLINE"
	STATUS_ONLINE_DpaClueProductListV2FilteringStatuses  DpaClueProductListV2FilteringStatuses = "STATUS_ONLINE"
)

// All allowed values of DpaClueProductListV2FilteringStatuses enum
var AllowedDpaClueProductListV2FilteringStatusesEnumValues = []DpaClueProductListV2FilteringStatuses{
	"STATUS_OFFLINE",
	"STATUS_ONLINE",
}

// NewDpaClueProductListV2FilteringStatusesFromValue returns a pointer to a valid DpaClueProductListV2FilteringStatuses
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaClueProductListV2FilteringStatusesFromValue(v string) (*DpaClueProductListV2FilteringStatuses, error) {
	ev := DpaClueProductListV2FilteringStatuses(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaClueProductListV2FilteringStatuses: valid values are %v", v, AllowedDpaClueProductListV2FilteringStatusesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaClueProductListV2FilteringStatuses) IsValid() bool {
	for _, existing := range AllowedDpaClueProductListV2FilteringStatusesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_clue_product_list_v2_filtering_statuses value
func (v DpaClueProductListV2FilteringStatuses) Ptr() *DpaClueProductListV2FilteringStatuses {
	return &v
}
