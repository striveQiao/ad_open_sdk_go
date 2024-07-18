/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanDmpAudiencesGetV10ResponseDataRetargetingTagsInner struct for QianchuanDmpAudiencesGetV10ResponseDataRetargetingTagsInner
type QianchuanDmpAudiencesGetV10ResponseDataRetargetingTagsInner struct {
	//
	CoverNum      *int64                                                       `json:"cover_num,omitempty"`
	HasOfflineTag *QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag `json:"has_offline_tag,omitempty"`
	IsCommon      *QianchuanDmpAudiencesGetV10DataRetargetingTagsIsCommon      `json:"is_common,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	RetargetingTagsId *int64                                                           `json:"retargeting_tags_id,omitempty"`
	RetargetingTagsOp *QianchuanDmpAudiencesGetV10DataRetargetingTagsRetargetingTagsOp `json:"retargeting_tags_op,omitempty"`
	//
	RetargetingTagsTip *string                                               `json:"retargeting_tags_tip,omitempty"`
	Source             *QianchuanDmpAudiencesGetV10DataRetargetingTagsSource `json:"source,omitempty"`
	//
	Status *int32 `json:"status,omitempty"`
}
