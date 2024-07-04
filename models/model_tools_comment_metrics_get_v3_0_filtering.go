/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentMetricsGetV30Filtering
type ToolsCommentMetricsGetV30Filtering struct {
	//
	AuthorIds  []int64                                       `json:"author_ids,omitempty"`
	HideStatus *ToolsCommentMetricsGetV30FilteringHideStatus `json:"hide_status,omitempty"`
	//
	ItemIds   []int64                                      `json:"item_ids,omitempty"`
	LevelType *ToolsCommentMetricsGetV30FilteringLevelType `json:"level_type,omitempty"`
}
