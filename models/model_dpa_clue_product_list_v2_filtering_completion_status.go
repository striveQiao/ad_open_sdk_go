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

// DpaClueProductListV2FilteringCompletionStatus
type DpaClueProductListV2FilteringCompletionStatus string

// List of dpa_clue_product_list_v2_filtering_completion_status
const (
	AD_COMPLETED_DpaClueProductListV2FilteringCompletionStatus    DpaClueProductListV2FilteringCompletionStatus = "AD_COMPLETED"
	ALL_COMPLETED_DpaClueProductListV2FilteringCompletionStatus   DpaClueProductListV2FilteringCompletionStatus = "ALL_COMPLETED"
	LEADS_COMPLETED_DpaClueProductListV2FilteringCompletionStatus DpaClueProductListV2FilteringCompletionStatus = "LEADS_COMPLETED"
	TO_BE_COMPLETED_DpaClueProductListV2FilteringCompletionStatus DpaClueProductListV2FilteringCompletionStatus = "TO_BE_COMPLETED"
)

// All allowed values of DpaClueProductListV2FilteringCompletionStatus enum
var AllowedDpaClueProductListV2FilteringCompletionStatusEnumValues = []DpaClueProductListV2FilteringCompletionStatus{
	"AD_COMPLETED",
	"ALL_COMPLETED",
	"LEADS_COMPLETED",
	"TO_BE_COMPLETED",
}

// NewDpaClueProductListV2FilteringCompletionStatusFromValue returns a pointer to a valid DpaClueProductListV2FilteringCompletionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaClueProductListV2FilteringCompletionStatusFromValue(v string) (*DpaClueProductListV2FilteringCompletionStatus, error) {
	ev := DpaClueProductListV2FilteringCompletionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaClueProductListV2FilteringCompletionStatus: valid values are %v", v, AllowedDpaClueProductListV2FilteringCompletionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaClueProductListV2FilteringCompletionStatus) IsValid() bool {
	for _, existing := range AllowedDpaClueProductListV2FilteringCompletionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_clue_product_list_v2_filtering_completion_status value
func (v DpaClueProductListV2FilteringCompletionStatus) Ptr() *DpaClueProductListV2FilteringCompletionStatus {
	return &v
}
