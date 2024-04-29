/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreativeListInnerAbstractMaterialsInnerTextAbstractInfo 文本摘要信息
type CreativeDetailGetV30ResponseDataCreativeListInnerAbstractMaterialsInnerTextAbstractInfo struct {
	// 文本摘要内容
	AbstractText *string `json:"abstract_text,omitempty"`
	// 搜索关键词列表
	BidwordList []*CreativeDetailGetV30ResponseDataCreativeListInnerAbstractMaterialsInnerTextAbstractInfoBidwordListInner `json:"bidword_list,omitempty"`
	// 动态词包ID列表
	WordList []*CreativeDetailGetV30ResponseDataCreativeListInnerAbstractMaterialsInnerTextAbstractInfoWordListInner `json:"word_list,omitempty"`
}
