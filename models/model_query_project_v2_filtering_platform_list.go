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

// QueryProjectV2FilteringPlatformList
type QueryProjectV2FilteringPlatformList string

// List of query_project_v2_filtering_platform_list
const (
	AD_QueryProjectV2FilteringPlatformList         QueryProjectV2FilteringPlatformList = "AD"
	NATIVELIFE_QueryProjectV2FilteringPlatformList QueryProjectV2FilteringPlatformList = "NATIVELIFE"
	QIANCHUAN_QueryProjectV2FilteringPlatformList  QueryProjectV2FilteringPlatformList = "QIANCHUAN"
)

// All allowed values of QueryProjectV2FilteringPlatformList enum
var AllowedQueryProjectV2FilteringPlatformListEnumValues = []QueryProjectV2FilteringPlatformList{
	"AD",
	"NATIVELIFE",
	"QIANCHUAN",
}

// NewQueryProjectV2FilteringPlatformListFromValue returns a pointer to a valid QueryProjectV2FilteringPlatformList
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryProjectV2FilteringPlatformListFromValue(v string) (*QueryProjectV2FilteringPlatformList, error) {
	ev := QueryProjectV2FilteringPlatformList(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryProjectV2FilteringPlatformList: valid values are %v", v, AllowedQueryProjectV2FilteringPlatformListEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryProjectV2FilteringPlatformList) IsValid() bool {
	for _, existing := range AllowedQueryProjectV2FilteringPlatformListEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_project_v2_filtering_platform_list value
func (v QueryProjectV2FilteringPlatformList) Ptr() *QueryProjectV2FilteringPlatformList {
	return &v
}
