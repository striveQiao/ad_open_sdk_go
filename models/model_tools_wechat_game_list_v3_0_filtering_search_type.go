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

// ToolsWechatGameListV30FilteringSearchType
type ToolsWechatGameListV30FilteringSearchType string

// List of tools_wechat_game_list_v3.0_filtering_search_type
const (
	CREATE_ONLY_ToolsWechatGameListV30FilteringSearchType ToolsWechatGameListV30FilteringSearchType = "CREATE_ONLY"
	SHARE_ONLY_ToolsWechatGameListV30FilteringSearchType  ToolsWechatGameListV30FilteringSearchType = "SHARE_ONLY"
)

// All allowed values of ToolsWechatGameListV30FilteringSearchType enum
var AllowedToolsWechatGameListV30FilteringSearchTypeEnumValues = []ToolsWechatGameListV30FilteringSearchType{
	"CREATE_ONLY",
	"SHARE_ONLY",
}

// NewToolsWechatGameListV30FilteringSearchTypeFromValue returns a pointer to a valid ToolsWechatGameListV30FilteringSearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatGameListV30FilteringSearchTypeFromValue(v string) (*ToolsWechatGameListV30FilteringSearchType, error) {
	ev := ToolsWechatGameListV30FilteringSearchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatGameListV30FilteringSearchType: valid values are %v", v, AllowedToolsWechatGameListV30FilteringSearchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatGameListV30FilteringSearchType) IsValid() bool {
	for _, existing := range AllowedToolsWechatGameListV30FilteringSearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_game_list_v3.0_filtering_search_type value
func (v ToolsWechatGameListV30FilteringSearchType) Ptr() *ToolsWechatGameListV30FilteringSearchType {
	return &v
}
