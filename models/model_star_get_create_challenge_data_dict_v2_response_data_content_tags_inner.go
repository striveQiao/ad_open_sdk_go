/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarGetCreateChallengeDataDictV2ResponseDataContentTagsInner struct for StarGetCreateChallengeDataDictV2ResponseDataContentTagsInner
type StarGetCreateChallengeDataDictV2ResponseDataContentTagsInner struct {
	// 父标签值，为0则无父标签
	ParentTagValue *int64 `json:"parent_tag_value,omitempty"`
	// 内容标签名称
	TagName *string `json:"tag_name,omitempty"`
	// 内容标签值
	TagValue *int64 `json:"tag_value,omitempty"`
}
