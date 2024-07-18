/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnGetAuthorListV2ResponseDataDataInner struct for StarMcnGetAuthorListV2ResponseDataDataInner
type StarMcnGetAuthorListV2ResponseDataDataInner struct {
	//
	AuthorId *int64 `json:"author_id,omitempty"`
	//
	AuthorName *string `json:"author_name,omitempty"`
	//
	IsBlock *bool `json:"is_block,omitempty"`
	//
	IsContracted *bool `json:"is_contracted,omitempty"`
	//
	IsRealName *bool `json:"is_real_name,omitempty"`
	//
	ValidFansAmount *int64 `json:"valid_fans_amount,omitempty"`
}
