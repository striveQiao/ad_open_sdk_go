/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2RequestCreativeTitleMaterialsInner struct for CreativeProceduralCreativeUpdateV2RequestCreativeTitleMaterialsInner
type CreativeProceduralCreativeUpdateV2RequestCreativeTitleMaterialsInner struct {
	//
	BidwordList []*CreativeProceduralCreativeUpdateV2RequestCreativeTitleMaterialsInnerBidwordListInner `json:"bidword_list,omitempty"`
	//
	DpaWordList []*CreativeProceduralCreativeUpdateV2RequestCreativeTitleMaterialsInnerDpaWordListInner `json:"dpa_word_list,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	Title string `json:"title"`
	//
	WordList []*CreativeProceduralCreativeUpdateV2RequestCreativeTitleMaterialsInnerWordListInner `json:"word_list,omitempty"`
}
