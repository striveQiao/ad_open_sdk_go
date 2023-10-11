/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderListV30FilterSceneType
type DouplusOrderListV30FilterSceneType string

// List of douplus_order_list_v3.0_filter_scene_type
const (
	LIVE_DouplusOrderListV30FilterSceneType  DouplusOrderListV30FilterSceneType = "LIVE"
	VIDEO_DouplusOrderListV30FilterSceneType DouplusOrderListV30FilterSceneType = "VIDEO"
)

// All allowed values of DouplusOrderListV30FilterSceneType enum
var AllowedDouplusOrderListV30FilterSceneTypeEnumValues = []DouplusOrderListV30FilterSceneType{
	"LIVE",
	"VIDEO",
}

// NewDouplusOrderListV30FilterSceneTypeFromValue returns a pointer to a valid DouplusOrderListV30FilterSceneType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30FilterSceneTypeFromValue(v string) (*DouplusOrderListV30FilterSceneType, error) {
	ev := DouplusOrderListV30FilterSceneType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30FilterSceneType: valid values are %v", v, AllowedDouplusOrderListV30FilterSceneTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30FilterSceneType) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30FilterSceneTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_filter_scene_type value
func (v DouplusOrderListV30FilterSceneType) Ptr() *DouplusOrderListV30FilterSceneType {
	return &v
}
