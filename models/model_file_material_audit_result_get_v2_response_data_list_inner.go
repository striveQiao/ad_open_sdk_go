/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialAuditResultGetV2ResponseDataListInner struct for FileMaterialAuditResultGetV2ResponseDataListInner
type FileMaterialAuditResultGetV2ResponseDataListInner struct {
	AuditStatus *FileMaterialAuditResultGetV2DataListAuditStatus `json:"audit_status,omitempty"`
	//
	MaterialId   *int64                                            `json:"material_id,omitempty"`
	MaterialType *FileMaterialAuditResultGetV2DataListMaterialType `json:"material_type,omitempty"`
	//
	Reasons []string `json:"reasons,omitempty"`
}
