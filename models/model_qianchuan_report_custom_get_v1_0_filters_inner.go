/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportCustomGetV10FiltersInner struct for QianchuanReportCustomGetV10FiltersInner
type QianchuanReportCustomGetV10FiltersInner struct {
	//
	Field string `json:"field"`
	//
	Operator int64 `json:"operator"`
	//
	Type *int64 `json:"type,omitempty"`
	//
	Values []string `json:"values"`
}
