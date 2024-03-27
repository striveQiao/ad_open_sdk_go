/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeEstimateProfitV10AudienceAge
type QianchuanAwemeEstimateProfitV10AudienceAge string

// List of qianchuan_aweme_estimate_profit_v1.0_audience_age
const (
	AGE_BETWEEN_18_23_QianchuanAwemeEstimateProfitV10AudienceAge QianchuanAwemeEstimateProfitV10AudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanAwemeEstimateProfitV10AudienceAge QianchuanAwemeEstimateProfitV10AudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanAwemeEstimateProfitV10AudienceAge QianchuanAwemeEstimateProfitV10AudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_50_QianchuanAwemeEstimateProfitV10AudienceAge QianchuanAwemeEstimateProfitV10AudienceAge = "AGE_BETWEEN_41_50"
)

// All allowed values of QianchuanAwemeEstimateProfitV10AudienceAge enum
var AllowedQianchuanAwemeEstimateProfitV10AudienceAgeEnumValues = []QianchuanAwemeEstimateProfitV10AudienceAge{
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_50",
}

// NewQianchuanAwemeEstimateProfitV10AudienceAgeFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10AudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10AudienceAgeFromValue(v string) (*QianchuanAwemeEstimateProfitV10AudienceAge, error) {
	ev := QianchuanAwemeEstimateProfitV10AudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10AudienceAge: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10AudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10AudienceAge) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10AudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_audience_age value
func (v QianchuanAwemeEstimateProfitV10AudienceAge) Ptr() *QianchuanAwemeEstimateProfitV10AudienceAge {
	return &v
}
