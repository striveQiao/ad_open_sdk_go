/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeStrategyListV2Response struct for CreativeStrategyListV2Response
type CreativeStrategyListV2Response struct {
	//
	Code *int64                              `json:"code,omitempty"`
	Data *CreativeStrategyListV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
