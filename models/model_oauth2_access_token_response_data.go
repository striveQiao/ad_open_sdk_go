/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// Oauth2AccessTokenResponseData
type Oauth2AccessTokenResponseData struct {
	//
	AccessToken *string `json:"access_token,omitempty"`
	//
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty"`
	//
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	//
	RefreshToken *string `json:"refresh_token,omitempty"`
	//
	RefreshTokenExpiresIn *int64 `json:"refresh_token_expires_in,omitempty"`
}
