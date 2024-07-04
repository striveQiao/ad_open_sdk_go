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

// QianchuanToolsEstimateAudienceV10AudienceNewCustomer
type QianchuanToolsEstimateAudienceV10AudienceNewCustomer string

// List of qianchuan_tools_estimate_audience_v1.0_audience_new_customer
const (
	NONE_QianchuanToolsEstimateAudienceV10AudienceNewCustomer   QianchuanToolsEstimateAudienceV10AudienceNewCustomer = "NONE"
	NO_BUY_QianchuanToolsEstimateAudienceV10AudienceNewCustomer QianchuanToolsEstimateAudienceV10AudienceNewCustomer = "NO_BUY"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceNewCustomer enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceNewCustomerEnumValues = []QianchuanToolsEstimateAudienceV10AudienceNewCustomer{
	"NONE",
	"NO_BUY",
}

// NewQianchuanToolsEstimateAudienceV10AudienceNewCustomerFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceNewCustomer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceNewCustomerFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceNewCustomer, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceNewCustomer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceNewCustomer: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceNewCustomerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceNewCustomer) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceNewCustomerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_new_customer value
func (v QianchuanToolsEstimateAudienceV10AudienceNewCustomer) Ptr() *QianchuanToolsEstimateAudienceV10AudienceNewCustomer {
	return &v
}
