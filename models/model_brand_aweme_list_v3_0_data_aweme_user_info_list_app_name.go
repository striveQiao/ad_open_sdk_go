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

// BrandAwemeListV30DataAwemeUserInfoListAppName
type BrandAwemeListV30DataAwemeUserInfoListAppName string

// List of brand_aweme_list_v3.0_data_aweme_user_info_list_app_name
const (
	AWEME_BrandAwemeListV30DataAwemeUserInfoListAppName      BrandAwemeListV30DataAwemeUserInfoListAppName = "AWEME"
	WATERMELON_BrandAwemeListV30DataAwemeUserInfoListAppName BrandAwemeListV30DataAwemeUserInfoListAppName = "WATERMELON"
)

// All allowed values of BrandAwemeListV30DataAwemeUserInfoListAppName enum
var AllowedBrandAwemeListV30DataAwemeUserInfoListAppNameEnumValues = []BrandAwemeListV30DataAwemeUserInfoListAppName{
	"AWEME",
	"WATERMELON",
}

// NewBrandAwemeListV30DataAwemeUserInfoListAppNameFromValue returns a pointer to a valid BrandAwemeListV30DataAwemeUserInfoListAppName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAwemeListV30DataAwemeUserInfoListAppNameFromValue(v string) (*BrandAwemeListV30DataAwemeUserInfoListAppName, error) {
	ev := BrandAwemeListV30DataAwemeUserInfoListAppName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAwemeListV30DataAwemeUserInfoListAppName: valid values are %v", v, AllowedBrandAwemeListV30DataAwemeUserInfoListAppNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAwemeListV30DataAwemeUserInfoListAppName) IsValid() bool {
	for _, existing := range AllowedBrandAwemeListV30DataAwemeUserInfoListAppNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_aweme_list_v3.0_data_aweme_user_info_list_app_name value
func (v BrandAwemeListV30DataAwemeUserInfoListAppName) Ptr() *BrandAwemeListV30DataAwemeUserInfoListAppName {
	return &v
}
