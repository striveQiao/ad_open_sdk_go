/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInner struct for ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInner
type ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInner struct {
	Scene *ToolsPromotionDiagnosisSuggestionGetV30DataSuggestionListSceneListScene `json:"scene,omitempty"`
	//
	Suggestions []*ToolsPromotionDiagnosisSuggestionGetV30ResponseDataSuggestionListInnerSceneListInnerSuggestionsInner `json:"suggestions,omitempty"`
}
