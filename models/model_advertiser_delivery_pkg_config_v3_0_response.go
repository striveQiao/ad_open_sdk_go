/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgConfigV30Response struct for AdvertiserDeliveryPkgConfigV30Response
type AdvertiserDeliveryPkgConfigV30Response struct {
	//
	Code *int64                                      `json:"code,omitempty"`
	Data *AdvertiserDeliveryPkgConfigV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
