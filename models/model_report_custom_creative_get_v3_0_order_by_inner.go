/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomCreativeGetV30OrderByInner struct for ReportCustomCreativeGetV30OrderByInner
type ReportCustomCreativeGetV30OrderByInner struct {
	// 排序字段
	Field string                                 `json:"field"`
	Type  *ReportCustomCreativeGetV30OrderByType `json:"type,omitempty"`
}
