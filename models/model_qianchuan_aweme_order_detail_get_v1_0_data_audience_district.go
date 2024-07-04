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

// QianchuanAwemeOrderDetailGetV10DataAudienceDistrict
type QianchuanAwemeOrderDetailGetV10DataAudienceDistrict string

// List of qianchuan_aweme_order_detail_get_v1.0_data_audience_district
const (
	CITY_QianchuanAwemeOrderDetailGetV10DataAudienceDistrict   QianchuanAwemeOrderDetailGetV10DataAudienceDistrict = "CITY"
	COUNTY_QianchuanAwemeOrderDetailGetV10DataAudienceDistrict QianchuanAwemeOrderDetailGetV10DataAudienceDistrict = "COUNTY"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataAudienceDistrict enum
var AllowedQianchuanAwemeOrderDetailGetV10DataAudienceDistrictEnumValues = []QianchuanAwemeOrderDetailGetV10DataAudienceDistrict{
	"CITY",
	"COUNTY",
}

// NewQianchuanAwemeOrderDetailGetV10DataAudienceDistrictFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataAudienceDistrictFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataAudienceDistrict, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataAudienceDistrict: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_audience_district value
func (v QianchuanAwemeOrderDetailGetV10DataAudienceDistrict) Ptr() *QianchuanAwemeOrderDetailGetV10DataAudienceDistrict {
	return &v
}
