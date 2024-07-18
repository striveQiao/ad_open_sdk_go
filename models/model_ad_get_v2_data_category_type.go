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

// AdGetV2DataCategoryType
type AdGetV2DataCategoryType string

// List of ad_get_v2_data_category_type
const (
	NONE_AdGetV2DataCategoryType     AdGetV2DataCategoryType = "NONE"
	CATEGORY_AdGetV2DataCategoryType AdGetV2DataCategoryType = "CATEGORY"
	PRODUCT_AdGetV2DataCategoryType  AdGetV2DataCategoryType = "PRODUCT"
)

// All allowed values of AdGetV2DataCategoryType enum
var AllowedAdGetV2DataCategoryTypeEnumValues = []AdGetV2DataCategoryType{
	"NONE",
	"CATEGORY",
	"PRODUCT",
}

// NewAdGetV2DataCategoryTypeFromValue returns a pointer to a valid AdGetV2DataCategoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataCategoryTypeFromValue(v string) (*AdGetV2DataCategoryType, error) {
	ev := AdGetV2DataCategoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataCategoryType: valid values are %v", v, AllowedAdGetV2DataCategoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataCategoryType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataCategoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_category_type value
func (v AdGetV2DataCategoryType) Ptr() *AdGetV2DataCategoryType {
	return &v
}
