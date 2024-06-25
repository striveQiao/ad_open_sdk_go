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

// EnterpriseCommentListGetV10FilterLevel
type EnterpriseCommentListGetV10FilterLevel string

// List of enterprise_comment_list_get_v1.0_filter_level
const (
	LEVEL_TWO_EnterpriseCommentListGetV10FilterLevel EnterpriseCommentListGetV10FilterLevel = "LEVEL_TWO"
	LEVEL_ALL_EnterpriseCommentListGetV10FilterLevel EnterpriseCommentListGetV10FilterLevel = "LEVEL_ALL"
	LEVEL_ONE_EnterpriseCommentListGetV10FilterLevel EnterpriseCommentListGetV10FilterLevel = "LEVEL_ONE"
)

// All allowed values of EnterpriseCommentListGetV10FilterLevel enum
var AllowedEnterpriseCommentListGetV10FilterLevelEnumValues = []EnterpriseCommentListGetV10FilterLevel{
	"LEVEL_TWO",
	"LEVEL_ALL",
	"LEVEL_ONE",
}

// NewEnterpriseCommentListGetV10FilterLevelFromValue returns a pointer to a valid EnterpriseCommentListGetV10FilterLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentListGetV10FilterLevelFromValue(v string) (*EnterpriseCommentListGetV10FilterLevel, error) {
	ev := EnterpriseCommentListGetV10FilterLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentListGetV10FilterLevel: valid values are %v", v, AllowedEnterpriseCommentListGetV10FilterLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentListGetV10FilterLevel) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentListGetV10FilterLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_list_get_v1.0_filter_level value
func (v EnterpriseCommentListGetV10FilterLevel) Ptr() *EnterpriseCommentListGetV10FilterLevel {
	return &v
}
