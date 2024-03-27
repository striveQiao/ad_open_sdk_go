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

// AdGetV2DataAutoUpdateKeyword
type AdGetV2DataAutoUpdateKeyword string

// List of ad_get_v2_data_auto_update_keyword
const (
	OFF_AdGetV2DataAutoUpdateKeyword AdGetV2DataAutoUpdateKeyword = "OFF"
	ON_AdGetV2DataAutoUpdateKeyword  AdGetV2DataAutoUpdateKeyword = "ON"
)

// All allowed values of AdGetV2DataAutoUpdateKeyword enum
var AllowedAdGetV2DataAutoUpdateKeywordEnumValues = []AdGetV2DataAutoUpdateKeyword{
	"OFF",
	"ON",
}

// NewAdGetV2DataAutoUpdateKeywordFromValue returns a pointer to a valid AdGetV2DataAutoUpdateKeyword
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAutoUpdateKeywordFromValue(v string) (*AdGetV2DataAutoUpdateKeyword, error) {
	ev := AdGetV2DataAutoUpdateKeyword(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAutoUpdateKeyword: valid values are %v", v, AllowedAdGetV2DataAutoUpdateKeywordEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAutoUpdateKeyword) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAutoUpdateKeywordEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_auto_update_keyword value
func (v AdGetV2DataAutoUpdateKeyword) Ptr() *AdGetV2DataAutoUpdateKeyword {
	return &v
}
