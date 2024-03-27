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

// EnterpriseCommentListGetV10OrderType
type EnterpriseCommentListGetV10OrderType string

// List of enterprise_comment_list_get_v1.0_order_type
const (
	DESC_EnterpriseCommentListGetV10OrderType EnterpriseCommentListGetV10OrderType = "DESC"
	ASC_EnterpriseCommentListGetV10OrderType  EnterpriseCommentListGetV10OrderType = "ASC"
)

// All allowed values of EnterpriseCommentListGetV10OrderType enum
var AllowedEnterpriseCommentListGetV10OrderTypeEnumValues = []EnterpriseCommentListGetV10OrderType{
	"DESC",
	"ASC",
}

// NewEnterpriseCommentListGetV10OrderTypeFromValue returns a pointer to a valid EnterpriseCommentListGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentListGetV10OrderTypeFromValue(v string) (*EnterpriseCommentListGetV10OrderType, error) {
	ev := EnterpriseCommentListGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentListGetV10OrderType: valid values are %v", v, AllowedEnterpriseCommentListGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentListGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentListGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_list_get_v1.0_order_type value
func (v EnterpriseCommentListGetV10OrderType) Ptr() *EnterpriseCommentListGetV10OrderType {
	return &v
}
