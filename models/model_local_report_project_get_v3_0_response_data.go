/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalReportProjectGetV30ResponseData
type LocalReportProjectGetV30ResponseData struct {
	PageInfo *LocalReportProjectGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	ProjectList []*LocalReportProjectGetV30ResponseDataProjectListInner `json:"project_list,omitempty"`
}
