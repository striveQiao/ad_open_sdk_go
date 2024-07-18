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

// QianchuanImageGetV10FilteringSources
type QianchuanImageGetV10FilteringSources string

// List of qianchuan_image_get_v1.0_filtering_sources
const (
	CREATIVE_CENTER_QianchuanImageGetV10FilteringSources QianchuanImageGetV10FilteringSources = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanImageGetV10FilteringSources      QianchuanImageGetV10FilteringSources = "E_COMMERCE"
	JI_CHUANG_QianchuanImageGetV10FilteringSources       QianchuanImageGetV10FilteringSources = "JI_CHUANG"
)

// All allowed values of QianchuanImageGetV10FilteringSources enum
var AllowedQianchuanImageGetV10FilteringSourcesEnumValues = []QianchuanImageGetV10FilteringSources{
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JI_CHUANG",
}

// NewQianchuanImageGetV10FilteringSourcesFromValue returns a pointer to a valid QianchuanImageGetV10FilteringSources
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanImageGetV10FilteringSourcesFromValue(v string) (*QianchuanImageGetV10FilteringSources, error) {
	ev := QianchuanImageGetV10FilteringSources(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanImageGetV10FilteringSources: valid values are %v", v, AllowedQianchuanImageGetV10FilteringSourcesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanImageGetV10FilteringSources) IsValid() bool {
	for _, existing := range AllowedQianchuanImageGetV10FilteringSourcesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_image_get_v1.0_filtering_sources value
func (v QianchuanImageGetV10FilteringSources) Ptr() *QianchuanImageGetV10FilteringSources {
	return &v
}
