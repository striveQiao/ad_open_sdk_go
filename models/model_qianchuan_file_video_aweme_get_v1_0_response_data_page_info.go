/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFileVideoAwemeGetV10ResponseDataPageInfo
type QianchuanFileVideoAwemeGetV10ResponseDataPageInfo struct {
	//
	Count *int64 `json:"count,omitempty"`
	//
	Cursor  *int64                                            `json:"cursor,omitempty"`
	HasMore *QianchuanFileVideoAwemeGetV10DataPageInfoHasMore `json:"has_more,omitempty"`
}
