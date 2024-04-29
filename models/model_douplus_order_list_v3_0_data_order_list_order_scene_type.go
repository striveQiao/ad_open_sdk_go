/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderListV30DataOrderListOrderSceneType
type DouplusOrderListV30DataOrderListOrderSceneType string

// List of douplus_order_list_v3.0_data_order_list_order_scene_type
const (
	LIVE_DouplusOrderListV30DataOrderListOrderSceneType  DouplusOrderListV30DataOrderListOrderSceneType = "LIVE"
	VIDEO_DouplusOrderListV30DataOrderListOrderSceneType DouplusOrderListV30DataOrderListOrderSceneType = "VIDEO"
)

// All allowed values of DouplusOrderListV30DataOrderListOrderSceneType enum
var AllowedDouplusOrderListV30DataOrderListOrderSceneTypeEnumValues = []DouplusOrderListV30DataOrderListOrderSceneType{
	"LIVE",
	"VIDEO",
}

// NewDouplusOrderListV30DataOrderListOrderSceneTypeFromValue returns a pointer to a valid DouplusOrderListV30DataOrderListOrderSceneType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30DataOrderListOrderSceneTypeFromValue(v string) (*DouplusOrderListV30DataOrderListOrderSceneType, error) {
	ev := DouplusOrderListV30DataOrderListOrderSceneType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30DataOrderListOrderSceneType: valid values are %v", v, AllowedDouplusOrderListV30DataOrderListOrderSceneTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30DataOrderListOrderSceneType) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30DataOrderListOrderSceneTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_data_order_list_order_scene_type value
func (v DouplusOrderListV30DataOrderListOrderSceneType) Ptr() *DouplusOrderListV30DataOrderListOrderSceneType {
	return &v
}
