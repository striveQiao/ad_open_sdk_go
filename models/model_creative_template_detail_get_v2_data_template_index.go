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

// CreativeTemplateDetailGetV2DataTemplateIndex
type CreativeTemplateDetailGetV2DataTemplateIndex string

// List of creative_template_detail_get_v2_data_template_index
const (
	BACK_CreativeTemplateDetailGetV2DataTemplateIndex  CreativeTemplateDetailGetV2DataTemplateIndex = "BACK"
	FRONT_CreativeTemplateDetailGetV2DataTemplateIndex CreativeTemplateDetailGetV2DataTemplateIndex = "FRONT"
)

// All allowed values of CreativeTemplateDetailGetV2DataTemplateIndex enum
var AllowedCreativeTemplateDetailGetV2DataTemplateIndexEnumValues = []CreativeTemplateDetailGetV2DataTemplateIndex{
	"BACK",
	"FRONT",
}

// NewCreativeTemplateDetailGetV2DataTemplateIndexFromValue returns a pointer to a valid CreativeTemplateDetailGetV2DataTemplateIndex
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTemplateDetailGetV2DataTemplateIndexFromValue(v string) (*CreativeTemplateDetailGetV2DataTemplateIndex, error) {
	ev := CreativeTemplateDetailGetV2DataTemplateIndex(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeTemplateDetailGetV2DataTemplateIndex: valid values are %v", v, AllowedCreativeTemplateDetailGetV2DataTemplateIndexEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeTemplateDetailGetV2DataTemplateIndex) IsValid() bool {
	for _, existing := range AllowedCreativeTemplateDetailGetV2DataTemplateIndexEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_template_detail_get_v2_data_template_index value
func (v CreativeTemplateDetailGetV2DataTemplateIndex) Ptr() *CreativeTemplateDetailGetV2DataTemplateIndex {
	return &v
}
