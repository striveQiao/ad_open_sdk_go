/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInner struct for ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInner
type ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInner struct {
	//
	PromotionId *int64 `json:"promotion_id,omitempty"`
	//
	SceneList []*ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInner `json:"scene_list,omitempty"`
}
