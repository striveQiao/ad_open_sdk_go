/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestCreativeInfoTitleMaterialsInner struct for AdlabGroupCreateV30RequestCreativeInfoTitleMaterialsInner
type AdlabGroupCreateV30RequestCreativeInfoTitleMaterialsInner struct {
	// 标题信息
	Title string `json:"title"`
	// 词包信息
	WordList []*AdlabGroupCreateV30RequestCreativeInfoTitleMaterialsInnerWordListInner `json:"word_list,omitempty"`
}
