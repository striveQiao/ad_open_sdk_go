/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserSelectV2ResponseDataCursorPageInfo
type AgentAdvertiserSelectV2ResponseDataCursorPageInfo struct {
	//
	Count *int64 `json:"count,omitempty"`
	//
	Cursor *int64 `json:"cursor,omitempty"`
	//
	HasMore *bool `json:"has_more,omitempty"`
	//
	TotalNumber *int64 `json:"total_number,omitempty"`
}
