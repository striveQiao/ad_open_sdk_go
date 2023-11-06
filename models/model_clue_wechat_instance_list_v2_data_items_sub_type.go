/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueWechatInstanceListV2DataItemsSubType
type ClueWechatInstanceListV2DataItemsSubType string

// List of clue_wechat_instance_list_v2_data_items_sub_type
const (
	NORMAL_ClueWechatInstanceListV2DataItemsSubType             ClueWechatInstanceListV2DataItemsSubType = "NORMAL"
	SINGLE_ClueWechatInstanceListV2DataItemsSubType             ClueWechatInstanceListV2DataItemsSubType = "SINGLE"
	ENTERPRISE_DEFAULT_ClueWechatInstanceListV2DataItemsSubType ClueWechatInstanceListV2DataItemsSubType = "ENTERPRISE_DEFAULT"
	ENTERPRISE_CODE_ClueWechatInstanceListV2DataItemsSubType    ClueWechatInstanceListV2DataItemsSubType = "ENTERPRISE_CODE"
)

// All allowed values of ClueWechatInstanceListV2DataItemsSubType enum
var AllowedClueWechatInstanceListV2DataItemsSubTypeEnumValues = []ClueWechatInstanceListV2DataItemsSubType{
	"NORMAL",
	"SINGLE",
	"ENTERPRISE_DEFAULT",
	"ENTERPRISE_CODE",
}

// NewClueWechatInstanceListV2DataItemsSubTypeFromValue returns a pointer to a valid ClueWechatInstanceListV2DataItemsSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatInstanceListV2DataItemsSubTypeFromValue(v string) (*ClueWechatInstanceListV2DataItemsSubType, error) {
	ev := ClueWechatInstanceListV2DataItemsSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatInstanceListV2DataItemsSubType: valid values are %v", v, AllowedClueWechatInstanceListV2DataItemsSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatInstanceListV2DataItemsSubType) IsValid() bool {
	for _, existing := range AllowedClueWechatInstanceListV2DataItemsSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_instance_list_v2_data_items_sub_type value
func (v ClueWechatInstanceListV2DataItemsSubType) Ptr() *ClueWechatInstanceListV2DataItemsSubType {
	return &v
}
