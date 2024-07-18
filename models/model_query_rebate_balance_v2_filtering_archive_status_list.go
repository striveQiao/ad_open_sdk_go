/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QueryRebateBalanceV2FilteringArchiveStatusList
type QueryRebateBalanceV2FilteringArchiveStatusList string

// List of query_rebate_balance_v2_filtering_archive_status_list
const (
	NO_RETRIEVE_QueryRebateBalanceV2FilteringArchiveStatusList QueryRebateBalanceV2FilteringArchiveStatusList = "NO_RETRIEVE"
	RETRIEVED_QueryRebateBalanceV2FilteringArchiveStatusList   QueryRebateBalanceV2FilteringArchiveStatusList = "RETRIEVED"
	ARCHIVED_QueryRebateBalanceV2FilteringArchiveStatusList    QueryRebateBalanceV2FilteringArchiveStatusList = "ARCHIVED"
)

// All allowed values of QueryRebateBalanceV2FilteringArchiveStatusList enum
var AllowedQueryRebateBalanceV2FilteringArchiveStatusListEnumValues = []QueryRebateBalanceV2FilteringArchiveStatusList{
	"NO_RETRIEVE",
	"RETRIEVED",
	"ARCHIVED",
}

// NewQueryRebateBalanceV2FilteringArchiveStatusListFromValue returns a pointer to a valid QueryRebateBalanceV2FilteringArchiveStatusList
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateBalanceV2FilteringArchiveStatusListFromValue(v string) (*QueryRebateBalanceV2FilteringArchiveStatusList, error) {
	ev := QueryRebateBalanceV2FilteringArchiveStatusList(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateBalanceV2FilteringArchiveStatusList: valid values are %v", v, AllowedQueryRebateBalanceV2FilteringArchiveStatusListEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateBalanceV2FilteringArchiveStatusList) IsValid() bool {
	for _, existing := range AllowedQueryRebateBalanceV2FilteringArchiveStatusListEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_balance_v2_filtering_archive_status_list value
func (v QueryRebateBalanceV2FilteringArchiveStatusList) Ptr() *QueryRebateBalanceV2FilteringArchiveStatusList {
	return &v
}
