/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeMaterialMetricsV2ResponseDataStatsValue struct for CreativeMaterialMetricsV2ResponseDataStatsValue
type CreativeMaterialMetricsV2ResponseDataStatsValue struct {
	// 派生素材信息
	DerivedMaterials []*CreativeMaterialMetricsV2ResponseDataStatsValueDerivedMaterialsInner `json:"derived_materials,omitempty"`
	// 源素材 ID
	MaterialId *int64 `json:"material_id,omitempty"`
	// 派生素材总计消耗（单位：元）
	TotalDerivedCost *float64 `json:"total_derived_cost,omitempty"`
}
