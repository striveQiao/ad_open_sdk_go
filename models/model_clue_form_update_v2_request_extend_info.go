/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormUpdateV2RequestExtendInfo
type ClueFormUpdateV2RequestExtendInfo struct {
	CountConfig  *ClueFormUpdateV2RequestExtendInfoCountConfig  `json:"count_config,omitempty"`
	SignUpConfig *ClueFormUpdateV2RequestExtendInfoSignUpConfig `json:"sign_up_config,omitempty"`
	//
	SuccessTip *string `json:"success_tip,omitempty"`
}
