/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomGetV30FiltersInner struct for ReportCustomGetV30FiltersInner
type ReportCustomGetV30FiltersInner struct {
	//
	Field string `json:"field"`
	//
	Operator int64 `json:"operator"`
	//
	Type int64 `json:"type"`
	//
	Values []string `json:"values"`
}
