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

// QianchuanCarouselGetV10OrderType
type QianchuanCarouselGetV10OrderType string

// List of qianchuan_carousel_get_v1.0_order_type
const (
	ASC_QianchuanCarouselGetV10OrderType  QianchuanCarouselGetV10OrderType = "ASC"
	DESC_QianchuanCarouselGetV10OrderType QianchuanCarouselGetV10OrderType = "DESC"
)

// All allowed values of QianchuanCarouselGetV10OrderType enum
var AllowedQianchuanCarouselGetV10OrderTypeEnumValues = []QianchuanCarouselGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanCarouselGetV10OrderTypeFromValue returns a pointer to a valid QianchuanCarouselGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCarouselGetV10OrderTypeFromValue(v string) (*QianchuanCarouselGetV10OrderType, error) {
	ev := QianchuanCarouselGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCarouselGetV10OrderType: valid values are %v", v, AllowedQianchuanCarouselGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCarouselGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanCarouselGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_carousel_get_v1.0_order_type value
func (v QianchuanCarouselGetV10OrderType) Ptr() *QianchuanCarouselGetV10OrderType {
	return &v
}
