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

// QueryRebateAccountingInfoV2FilteringIsCreateRebate
type QueryRebateAccountingInfoV2FilteringIsCreateRebate string

// List of query_rebate_accounting_info_v2_filtering_is_create_rebate
const (
	YES_QueryRebateAccountingInfoV2FilteringIsCreateRebate QueryRebateAccountingInfoV2FilteringIsCreateRebate = "YES"
	NO_QueryRebateAccountingInfoV2FilteringIsCreateRebate  QueryRebateAccountingInfoV2FilteringIsCreateRebate = "NO"
)

// All allowed values of QueryRebateAccountingInfoV2FilteringIsCreateRebate enum
var AllowedQueryRebateAccountingInfoV2FilteringIsCreateRebateEnumValues = []QueryRebateAccountingInfoV2FilteringIsCreateRebate{
	"YES",
	"NO",
}

// NewQueryRebateAccountingInfoV2FilteringIsCreateRebateFromValue returns a pointer to a valid QueryRebateAccountingInfoV2FilteringIsCreateRebate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateAccountingInfoV2FilteringIsCreateRebateFromValue(v string) (*QueryRebateAccountingInfoV2FilteringIsCreateRebate, error) {
	ev := QueryRebateAccountingInfoV2FilteringIsCreateRebate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateAccountingInfoV2FilteringIsCreateRebate: valid values are %v", v, AllowedQueryRebateAccountingInfoV2FilteringIsCreateRebateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateAccountingInfoV2FilteringIsCreateRebate) IsValid() bool {
	for _, existing := range AllowedQueryRebateAccountingInfoV2FilteringIsCreateRebateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_accounting_info_v2_filtering_is_create_rebate value
func (v QueryRebateAccountingInfoV2FilteringIsCreateRebate) Ptr() *QueryRebateAccountingInfoV2FilteringIsCreateRebate {
	return &v
}
