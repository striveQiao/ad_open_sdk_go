/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2AdvertiserGetResponse struct for Oauth2AdvertiserGetResponse
type Oauth2AdvertiserGetResponse struct {
	//
	Code *int64                           `json:"code,omitempty"`
	Data *Oauth2AdvertiserGetResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
