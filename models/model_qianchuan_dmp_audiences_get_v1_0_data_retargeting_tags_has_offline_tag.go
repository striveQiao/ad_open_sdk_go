/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag
type QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag int64

// List of qianchuan_dmp_audiences_get_v1.0_data_retargeting_tags_has_offline_tag
const (
	Enum_0_QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag = 0
	Enum_1_QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag = 1
)

// All allowed values of QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag enum
var AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTagEnumValues = []QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag{
	0,
	1,
}

// NewQianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTagFromValue returns a pointer to a valid QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTagFromValue(v int64) (*QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag, error) {
	ev := QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag: valid values are %v", v, AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag) IsValid() bool {
	for _, existing := range AllowedQianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_dmp_audiences_get_v1.0_data_retargeting_tags_has_offline_tag value
func (v QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag) Ptr() *QianchuanDmpAudiencesGetV10DataRetargetingTagsHasOfflineTag {
	return &v
}
