/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCarouselAwemeGetV10DataPageInfoHasMore
type QianchuanCarouselAwemeGetV10DataPageInfoHasMore int64

// List of qianchuan_carousel_aweme_get_v1.0_data_page_info_has_more
const (
	Enum_1_QianchuanCarouselAwemeGetV10DataPageInfoHasMore QianchuanCarouselAwemeGetV10DataPageInfoHasMore = 1
	Enum_0_QianchuanCarouselAwemeGetV10DataPageInfoHasMore QianchuanCarouselAwemeGetV10DataPageInfoHasMore = 0
)

// All allowed values of QianchuanCarouselAwemeGetV10DataPageInfoHasMore enum
var AllowedQianchuanCarouselAwemeGetV10DataPageInfoHasMoreEnumValues = []QianchuanCarouselAwemeGetV10DataPageInfoHasMore{
	1,
	0,
}

// NewQianchuanCarouselAwemeGetV10DataPageInfoHasMoreFromValue returns a pointer to a valid QianchuanCarouselAwemeGetV10DataPageInfoHasMore
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCarouselAwemeGetV10DataPageInfoHasMoreFromValue(v int64) (*QianchuanCarouselAwemeGetV10DataPageInfoHasMore, error) {
	ev := QianchuanCarouselAwemeGetV10DataPageInfoHasMore(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCarouselAwemeGetV10DataPageInfoHasMore: valid values are %v", v, AllowedQianchuanCarouselAwemeGetV10DataPageInfoHasMoreEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCarouselAwemeGetV10DataPageInfoHasMore) IsValid() bool {
	for _, existing := range AllowedQianchuanCarouselAwemeGetV10DataPageInfoHasMoreEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_carousel_aweme_get_v1.0_data_page_info_has_more value
func (v QianchuanCarouselAwemeGetV10DataPageInfoHasMore) Ptr() *QianchuanCarouselAwemeGetV10DataPageInfoHasMore {
	return &v
}
