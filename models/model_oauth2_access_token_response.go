/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2AccessTokenResponse struct for Oauth2AccessTokenResponse
type Oauth2AccessTokenResponse struct {
	//
	Code *int64                         `json:"code,omitempty"`
	Data *Oauth2AccessTokenResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
