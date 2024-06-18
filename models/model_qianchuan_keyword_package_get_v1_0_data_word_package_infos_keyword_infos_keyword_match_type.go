/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType
type QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType string

// List of qianchuan_keyword_package_get_v1.0_data_word_package_infos_keyword_infos_keyword_match_type
const (
	EXTENSIVE_QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType = "EXTENSIVE"
	PHRASE_QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType    QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType = "PHRASE"
	PRECISION_QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType = "PRECISION"
)

// All allowed values of QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType enum
var AllowedQianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchTypeEnumValues = []QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewQianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchTypeFromValue returns a pointer to a valid QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchTypeFromValue(v string) (*QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType, error) {
	ev := QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType: valid values are %v", v, AllowedQianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType) IsValid() bool {
	for _, existing := range AllowedQianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_keyword_package_get_v1.0_data_word_package_infos_keyword_infos_keyword_match_type value
func (v QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType) Ptr() *QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType {
	return &v
}
