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

// EnterpriseFlowCategoryGetV10TimeDimension
type EnterpriseFlowCategoryGetV10TimeDimension string

// List of enterprise_flow_category_get_v1.0_time_dimension
const (
	SUM_EnterpriseFlowCategoryGetV10TimeDimension   EnterpriseFlowCategoryGetV10TimeDimension = "SUM"
	DAILY_EnterpriseFlowCategoryGetV10TimeDimension EnterpriseFlowCategoryGetV10TimeDimension = "DAILY"
)

// All allowed values of EnterpriseFlowCategoryGetV10TimeDimension enum
var AllowedEnterpriseFlowCategoryGetV10TimeDimensionEnumValues = []EnterpriseFlowCategoryGetV10TimeDimension{
	"SUM",
	"DAILY",
}

// NewEnterpriseFlowCategoryGetV10TimeDimensionFromValue returns a pointer to a valid EnterpriseFlowCategoryGetV10TimeDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseFlowCategoryGetV10TimeDimensionFromValue(v string) (*EnterpriseFlowCategoryGetV10TimeDimension, error) {
	ev := EnterpriseFlowCategoryGetV10TimeDimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseFlowCategoryGetV10TimeDimension: valid values are %v", v, AllowedEnterpriseFlowCategoryGetV10TimeDimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseFlowCategoryGetV10TimeDimension) IsValid() bool {
	for _, existing := range AllowedEnterpriseFlowCategoryGetV10TimeDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_flow_category_get_v1.0_time_dimension value
func (v EnterpriseFlowCategoryGetV10TimeDimension) Ptr() *EnterpriseFlowCategoryGetV10TimeDimension {
	return &v
}
