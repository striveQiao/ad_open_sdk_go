/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2AccessTokenRequest struct for Oauth2AccessTokenRequest
type Oauth2AccessTokenRequest struct {
	//
	AppId *int64 `json:"app_id,omitempty"`
	//
	AuthCode string `json:"auth_code"`
	//
	Secret string `json:"secret"`
}
