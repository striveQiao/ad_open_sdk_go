/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolQuickAppManagementQuickAppGetV2Response struct for ToolQuickAppManagementQuickAppGetV2Response
type ToolQuickAppManagementQuickAppGetV2Response struct {
	//
	Code *int64                                           `json:"code,omitempty"`
	Data *ToolQuickAppManagementQuickAppGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
