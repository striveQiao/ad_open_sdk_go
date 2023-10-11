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

// QianchuanAwemeEstimateProfitV10AudienceDistrict
type QianchuanAwemeEstimateProfitV10AudienceDistrict string

// List of qianchuan_aweme_estimate_profit_v1.0_audience_district
const (
	CITY_QianchuanAwemeEstimateProfitV10AudienceDistrict   QianchuanAwemeEstimateProfitV10AudienceDistrict = "CITY"
	COUNTY_QianchuanAwemeEstimateProfitV10AudienceDistrict QianchuanAwemeEstimateProfitV10AudienceDistrict = "COUNTY"
)

// All allowed values of QianchuanAwemeEstimateProfitV10AudienceDistrict enum
var AllowedQianchuanAwemeEstimateProfitV10AudienceDistrictEnumValues = []QianchuanAwemeEstimateProfitV10AudienceDistrict{
	"CITY",
	"COUNTY",
}

// NewQianchuanAwemeEstimateProfitV10AudienceDistrictFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10AudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10AudienceDistrictFromValue(v string) (*QianchuanAwemeEstimateProfitV10AudienceDistrict, error) {
	ev := QianchuanAwemeEstimateProfitV10AudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10AudienceDistrict: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10AudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10AudienceDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10AudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_audience_district value
func (v QianchuanAwemeEstimateProfitV10AudienceDistrict) Ptr() *QianchuanAwemeEstimateProfitV10AudienceDistrict {
	return &v
}
