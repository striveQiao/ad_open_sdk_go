/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueFormDetailV2ResponseDataElementsInner struct for ToolsClueFormDetailV2ResponseDataElementsInner
type ToolsClueFormDetailV2ResponseDataElementsInner struct {
	AllowEmpty *ToolsClueFormDetailV2DataElementsAllowEmpty `json:"allow_empty,omitempty"`
	IsUnique   *ToolsClueFormDetailV2DataElementsIsUnique   `json:"is_unique,omitempty"`
	//
	Label *string                                `json:"label,omitempty"`
	Type  *ToolsClueFormDetailV2DataElementsType `json:"type,omitempty"`
}
