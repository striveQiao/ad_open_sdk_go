/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerShareGetV30ResponseData
type EventManagerShareGetV30ResponseData struct {
	//
	List     []*EventManagerShareGetV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *EventManagerShareGetV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
