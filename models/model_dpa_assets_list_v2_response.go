/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetsListV2Response struct for DpaAssetsListV2Response
type DpaAssetsListV2Response struct {
	//
	Code *int64                       `json:"code,omitempty"`
	Data *DpaAssetsListV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
