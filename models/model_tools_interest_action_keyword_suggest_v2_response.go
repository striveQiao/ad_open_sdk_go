/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsInterestActionKeywordSuggestV2Response struct for ToolsInterestActionKeywordSuggestV2Response
type ToolsInterestActionKeywordSuggestV2Response struct {
	//
	Code *int64                                           `json:"code,omitempty"`
	Data *ToolsInterestActionKeywordSuggestV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
