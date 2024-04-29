/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QueryProjectV2FilteringReceiptStatusList
type QueryProjectV2FilteringReceiptStatusList string

// List of query_project_v2_filtering_receipt_status_list
const (
	COMPLETED_QueryProjectV2FilteringReceiptStatusList      QueryProjectV2FilteringReceiptStatusList = "COMPLETED"
	PART_COMPLETED_QueryProjectV2FilteringReceiptStatusList QueryProjectV2FilteringReceiptStatusList = "PART_COMPLETED"
	UNCOMPLETED_QueryProjectV2FilteringReceiptStatusList    QueryProjectV2FilteringReceiptStatusList = "UNCOMPLETED"
)

// All allowed values of QueryProjectV2FilteringReceiptStatusList enum
var AllowedQueryProjectV2FilteringReceiptStatusListEnumValues = []QueryProjectV2FilteringReceiptStatusList{
	"COMPLETED",
	"PART_COMPLETED",
	"UNCOMPLETED",
}

// NewQueryProjectV2FilteringReceiptStatusListFromValue returns a pointer to a valid QueryProjectV2FilteringReceiptStatusList
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryProjectV2FilteringReceiptStatusListFromValue(v string) (*QueryProjectV2FilteringReceiptStatusList, error) {
	ev := QueryProjectV2FilteringReceiptStatusList(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryProjectV2FilteringReceiptStatusList: valid values are %v", v, AllowedQueryProjectV2FilteringReceiptStatusListEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryProjectV2FilteringReceiptStatusList) IsValid() bool {
	for _, existing := range AllowedQueryProjectV2FilteringReceiptStatusListEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_project_v2_filtering_receipt_status_list value
func (v QueryProjectV2FilteringReceiptStatusList) Ptr() *QueryProjectV2FilteringReceiptStatusList {
	return &v
}
