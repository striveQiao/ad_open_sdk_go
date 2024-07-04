/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceDistrict
type QianchuanAdCreateV10AudienceDistrict string

// List of qianchuan_ad_create_v1.0_audience_district
const (
	CITY_QianchuanAdCreateV10AudienceDistrict   QianchuanAdCreateV10AudienceDistrict = "CITY"
	COUNTY_QianchuanAdCreateV10AudienceDistrict QianchuanAdCreateV10AudienceDistrict = "COUNTY"
	NONE_QianchuanAdCreateV10AudienceDistrict   QianchuanAdCreateV10AudienceDistrict = "NONE"
)

// All allowed values of QianchuanAdCreateV10AudienceDistrict enum
var AllowedQianchuanAdCreateV10AudienceDistrictEnumValues = []QianchuanAdCreateV10AudienceDistrict{
	"CITY",
	"COUNTY",
	"NONE",
}

// NewQianchuanAdCreateV10AudienceDistrictFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceDistrictFromValue(v string) (*QianchuanAdCreateV10AudienceDistrict, error) {
	ev := QianchuanAdCreateV10AudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceDistrict: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_district value
func (v QianchuanAdCreateV10AudienceDistrict) Ptr() *QianchuanAdCreateV10AudienceDistrict {
	return &v
}
