/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAgentGetV2V2ResponseDataCursorInfo
type ReportAgentGetV2V2ResponseDataCursorInfo struct {
	//
	Cursor *int64 `json:"cursor,omitempty"`
	//
	CursorSize *int64 `json:"cursor_size,omitempty"`
	//
	HasMore *int64 `json:"has_more,omitempty"`
	//
	TotalNumber *int64 `json:"total_number,omitempty"`
}
