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

// ClueWechatInstanceDetailV2DataSuffix
type ClueWechatInstanceDetailV2DataSuffix string

// List of clue_wechat_instance_detail_v2_data_suffix
const (
	FALSE_ClueWechatInstanceDetailV2DataSuffix ClueWechatInstanceDetailV2DataSuffix = "FALSE"
	TRUE_ClueWechatInstanceDetailV2DataSuffix  ClueWechatInstanceDetailV2DataSuffix = "TRUE"
)

// All allowed values of ClueWechatInstanceDetailV2DataSuffix enum
var AllowedClueWechatInstanceDetailV2DataSuffixEnumValues = []ClueWechatInstanceDetailV2DataSuffix{
	"FALSE",
	"TRUE",
}

// NewClueWechatInstanceDetailV2DataSuffixFromValue returns a pointer to a valid ClueWechatInstanceDetailV2DataSuffix
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatInstanceDetailV2DataSuffixFromValue(v string) (*ClueWechatInstanceDetailV2DataSuffix, error) {
	ev := ClueWechatInstanceDetailV2DataSuffix(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatInstanceDetailV2DataSuffix: valid values are %v", v, AllowedClueWechatInstanceDetailV2DataSuffixEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatInstanceDetailV2DataSuffix) IsValid() bool {
	for _, existing := range AllowedClueWechatInstanceDetailV2DataSuffixEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_instance_detail_v2_data_suffix value
func (v ClueWechatInstanceDetailV2DataSuffix) Ptr() *ClueWechatInstanceDetailV2DataSuffix {
	return &v
}
