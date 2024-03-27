/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QueryProjectV2FilteringServingTypeList
type QueryProjectV2FilteringServingTypeList string

// List of query_project_v2_filtering_serving_type_list
const (
	PINPAI_QueryProjectV2FilteringServingTypeList      QueryProjectV2FilteringServingTypeList = "PINPAI"
	PINPAIPMP_QueryProjectV2FilteringServingTypeList   QueryProjectV2FilteringServingTypeList = "PINPAIPMP"
	XIAOGUO_QueryProjectV2FilteringServingTypeList     QueryProjectV2FilteringServingTypeList = "XIAOGUO"
	XIAOGUO_RTB_QueryProjectV2FilteringServingTypeList QueryProjectV2FilteringServingTypeList = "XIAOGUO_RTB"
)

// All allowed values of QueryProjectV2FilteringServingTypeList enum
var AllowedQueryProjectV2FilteringServingTypeListEnumValues = []QueryProjectV2FilteringServingTypeList{
	"PINPAI",
	"PINPAIPMP",
	"XIAOGUO",
	"XIAOGUO_RTB",
}

// NewQueryProjectV2FilteringServingTypeListFromValue returns a pointer to a valid QueryProjectV2FilteringServingTypeList
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryProjectV2FilteringServingTypeListFromValue(v string) (*QueryProjectV2FilteringServingTypeList, error) {
	ev := QueryProjectV2FilteringServingTypeList(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryProjectV2FilteringServingTypeList: valid values are %v", v, AllowedQueryProjectV2FilteringServingTypeListEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryProjectV2FilteringServingTypeList) IsValid() bool {
	for _, existing := range AllowedQueryProjectV2FilteringServingTypeListEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_project_v2_filtering_serving_type_list value
func (v QueryProjectV2FilteringServingTypeList) Ptr() *QueryProjectV2FilteringServingTypeList {
	return &v
}
