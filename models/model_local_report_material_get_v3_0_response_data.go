/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalReportMaterialGetV30ResponseData
type LocalReportMaterialGetV30ResponseData struct {
	//
	MaterialList []*LocalReportMaterialGetV30ResponseDataMaterialListInner `json:"material_list,omitempty"`
	PageInfo     *LocalReportMaterialGetV30ResponseDataPageInfo            `json:"page_info,omitempty"`
}
