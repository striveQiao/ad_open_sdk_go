/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaTemplateGetV2ResponseData
type DpaTemplateGetV2ResponseData struct {
	//
	List     []*DpaTemplateGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *DpaTemplateGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
