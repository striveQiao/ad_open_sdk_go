/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomConfigGetV30ResponseDataListInnerDimensionsInnerFilterConfig
type ReportCustomConfigGetV30ResponseDataListInnerDimensionsInnerFilterConfig struct {
	//
	Operator *int64 `json:"operator,omitempty"`
	//
	RangeValue []*ReportCustomConfigGetV30ResponseDataListInnerDimensionsInnerFilterConfigRangeValueInner `json:"range_value,omitempty"`
	//
	Type *int64 `json:"type,omitempty"`
	//
	ValueLimit *int64 `json:"value_limit,omitempty"`
}
