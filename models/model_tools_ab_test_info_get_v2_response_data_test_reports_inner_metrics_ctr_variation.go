/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestInfoGetV2ResponseDataTestReportsInnerMetricsCtrVariation ctr波动范围
type ToolsAbTestInfoGetV2ResponseDataTestReportsInnerMetricsCtrVariation struct {
	// 波动范围上限
	VariationMax *string `json:"variation_max,omitempty"`
	// 波动范围下限
	VariationMin *string `json:"variation_min,omitempty"`
}
