/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormCreateV2RequestExtendInfo
type ClueFormCreateV2RequestExtendInfo struct {
	CountConfig  *ClueFormCreateV2RequestExtendInfoCountConfig  `json:"count_config,omitempty"`
	SignUpConfig *ClueFormCreateV2RequestExtendInfoSignUpConfig `json:"sign_up_config,omitempty"`
	//
	SuccessTip *string `json:"success_tip,omitempty"`
}
