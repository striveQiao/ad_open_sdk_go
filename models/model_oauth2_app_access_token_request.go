/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2AppAccessTokenRequest struct for Oauth2AppAccessTokenRequest
type Oauth2AppAccessTokenRequest struct {
	// 应用ID
	AppId int64 `json:"app_id"`
	// 应用秘钥
	Secret string `json:"secret"`
}
