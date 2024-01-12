/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeSuggestBidV10AudienceDistrict
type QianchuanAwemeSuggestBidV10AudienceDistrict string

// List of qianchuan_aweme_suggest_bid_v1.0_audience_district
const (
	CITY_QianchuanAwemeSuggestBidV10AudienceDistrict   QianchuanAwemeSuggestBidV10AudienceDistrict = "CITY"
	COUNTY_QianchuanAwemeSuggestBidV10AudienceDistrict QianchuanAwemeSuggestBidV10AudienceDistrict = "COUNTY"
)

// All allowed values of QianchuanAwemeSuggestBidV10AudienceDistrict enum
var AllowedQianchuanAwemeSuggestBidV10AudienceDistrictEnumValues = []QianchuanAwemeSuggestBidV10AudienceDistrict{
	"CITY",
	"COUNTY",
}

// NewQianchuanAwemeSuggestBidV10AudienceDistrictFromValue returns a pointer to a valid QianchuanAwemeSuggestBidV10AudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeSuggestBidV10AudienceDistrictFromValue(v string) (*QianchuanAwemeSuggestBidV10AudienceDistrict, error) {
	ev := QianchuanAwemeSuggestBidV10AudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeSuggestBidV10AudienceDistrict: valid values are %v", v, AllowedQianchuanAwemeSuggestBidV10AudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeSuggestBidV10AudienceDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeSuggestBidV10AudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_suggest_bid_v1.0_audience_district value
func (v QianchuanAwemeSuggestBidV10AudienceDistrict) Ptr() *QianchuanAwemeSuggestBidV10AudienceDistrict {
	return &v
}
