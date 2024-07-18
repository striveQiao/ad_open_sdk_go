/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// LocalReportProjectGetV30OrderType
type LocalReportProjectGetV30OrderType string

// List of local_report_project_get_v3.0_order_type
const (
	ASC_LocalReportProjectGetV30OrderType  LocalReportProjectGetV30OrderType = "ASC"
	DESC_LocalReportProjectGetV30OrderType LocalReportProjectGetV30OrderType = "DESC"
)

// All allowed values of LocalReportProjectGetV30OrderType enum
var AllowedLocalReportProjectGetV30OrderTypeEnumValues = []LocalReportProjectGetV30OrderType{
	"ASC",
	"DESC",
}

// NewLocalReportProjectGetV30OrderTypeFromValue returns a pointer to a valid LocalReportProjectGetV30OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalReportProjectGetV30OrderTypeFromValue(v string) (*LocalReportProjectGetV30OrderType, error) {
	ev := LocalReportProjectGetV30OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalReportProjectGetV30OrderType: valid values are %v", v, AllowedLocalReportProjectGetV30OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalReportProjectGetV30OrderType) IsValid() bool {
	for _, existing := range AllowedLocalReportProjectGetV30OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to local_report_project_get_v3.0_order_type value
func (v LocalReportProjectGetV30OrderType) Ptr() *LocalReportProjectGetV30OrderType {
	return &v
}
