/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType
type QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType string

// List of qianchuan_orientation_package_get_v1.0_data_list_InActive_retargeting_tags_InActive_type
const (
	EXPIRE_QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType         QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType = "EXPIRE"
	TAG_OFFLINE_QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType    QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType = "TAG_OFFLINE"
	MANUAL_OFFLINE_QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType = "MANUAL_OFFLINE"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType enum
var AllowedQianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveTypeEnumValues = []QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType{
	"EXPIRE",
	"TAG_OFFLINE",
	"MANUAL_OFFLINE",
}

// NewQianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveTypeFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveTypeFromValue(v string) (*QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType, error) {
	ev := QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_InActive_retargeting_tags_InActive_type value
func (v QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType) Ptr() *QianchuanOrientationPackageGetV10DataListInActiveRetargetingTagsInActiveType {
	return &v
}
