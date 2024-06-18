/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30RelatedProductProductSetting
type ProjectCreateV30RelatedProductProductSetting string

// List of project_create_v3.0_related_product_product_setting
const (
	NO_MAP_ProjectCreateV30RelatedProductProductSetting ProjectCreateV30RelatedProductProductSetting = "NO_MAP"
	SINGLE_ProjectCreateV30RelatedProductProductSetting ProjectCreateV30RelatedProductProductSetting = "SINGLE"
)

// All allowed values of ProjectCreateV30RelatedProductProductSetting enum
var AllowedProjectCreateV30RelatedProductProductSettingEnumValues = []ProjectCreateV30RelatedProductProductSetting{
	"NO_MAP",
	"SINGLE",
}

// NewProjectCreateV30RelatedProductProductSettingFromValue returns a pointer to a valid ProjectCreateV30RelatedProductProductSetting
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30RelatedProductProductSettingFromValue(v string) (*ProjectCreateV30RelatedProductProductSetting, error) {
	ev := ProjectCreateV30RelatedProductProductSetting(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30RelatedProductProductSetting: valid values are %v", v, AllowedProjectCreateV30RelatedProductProductSettingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30RelatedProductProductSetting) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30RelatedProductProductSettingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_related_product_product_setting value
func (v ProjectCreateV30RelatedProductProductSetting) Ptr() *ProjectCreateV30RelatedProductProductSetting {
	return &v
}
