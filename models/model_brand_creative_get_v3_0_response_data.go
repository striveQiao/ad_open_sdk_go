/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseData
type BrandCreativeGetV30ResponseData struct {
	//
	Creatives []*BrandCreativeGetV30ResponseDataCreativesInner `json:"creatives,omitempty"`
	// 结果总数
	TotalCount *int64 `json:"total_count,omitempty"`
}
