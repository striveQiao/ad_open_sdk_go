/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteFormsListV2Response struct for ToolsSiteFormsListV2Response
type ToolsSiteFormsListV2Response struct {
	//
	Code *int64                            `json:"code,omitempty"`
	Data *ToolsSiteFormsListV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
