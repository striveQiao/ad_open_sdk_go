/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AsyncTaskGetV2ResponseData
type AsyncTaskGetV2ResponseData struct {
	//
	List     []map[string]interface{}            `json:"list,omitempty"`
	PageInfo *AsyncTaskGetV2ResponseDataPageInfo `json:"page_info,omitempty"`
}
