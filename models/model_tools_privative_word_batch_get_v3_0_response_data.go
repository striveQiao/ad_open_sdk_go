/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordBatchGetV30ResponseData
type ToolsPrivativeWordBatchGetV30ResponseData struct {
	//
	ProjectsPrivative []*ToolsPrivativeWordBatchGetV30ResponseDataProjectsPrivativeInner `json:"projects_privative,omitempty"`
	//
	PromotionsPrivative []*ToolsPrivativeWordBatchGetV30ResponseDataPromotionsPrivativeInner `json:"promotions_privative,omitempty"`
}
