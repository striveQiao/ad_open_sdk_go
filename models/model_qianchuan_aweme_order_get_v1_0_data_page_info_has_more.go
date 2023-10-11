/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderGetV10DataPageInfoHasMore
type QianchuanAwemeOrderGetV10DataPageInfoHasMore int64

// List of qianchuan_aweme_order_get_v1.0_data_page_info_has_more
const (
	Enum_0_QianchuanAwemeOrderGetV10DataPageInfoHasMore QianchuanAwemeOrderGetV10DataPageInfoHasMore = 0
	Enum_1_QianchuanAwemeOrderGetV10DataPageInfoHasMore QianchuanAwemeOrderGetV10DataPageInfoHasMore = 1
)

// All allowed values of QianchuanAwemeOrderGetV10DataPageInfoHasMore enum
var AllowedQianchuanAwemeOrderGetV10DataPageInfoHasMoreEnumValues = []QianchuanAwemeOrderGetV10DataPageInfoHasMore{
	0,
	1,
}

// NewQianchuanAwemeOrderGetV10DataPageInfoHasMoreFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10DataPageInfoHasMore
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10DataPageInfoHasMoreFromValue(v int64) (*QianchuanAwemeOrderGetV10DataPageInfoHasMore, error) {
	ev := QianchuanAwemeOrderGetV10DataPageInfoHasMore(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10DataPageInfoHasMore: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10DataPageInfoHasMoreEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10DataPageInfoHasMore) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10DataPageInfoHasMoreEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_data_page_info_has_more value
func (v QianchuanAwemeOrderGetV10DataPageInfoHasMore) Ptr() *QianchuanAwemeOrderGetV10DataPageInfoHasMore {
	return &v
}
