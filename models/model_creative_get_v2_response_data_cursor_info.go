/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeGetV2ResponseDataCursorInfo
type CreativeGetV2ResponseDataCursorInfo struct {
	//
	Count *int64 `json:"count,omitempty"`
	//
	Cursor *int64 `json:"cursor,omitempty"`
	//
	HasMore *bool `json:"has_more,omitempty"`
}
