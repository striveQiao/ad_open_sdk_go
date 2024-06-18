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

// ToolsPreAuditSendV2PreAuditMaterialsType
type ToolsPreAuditSendV2PreAuditMaterialsType string

// List of tools_pre_audit_send_v2_pre_audit_materials_type
const (
	IMG_ToolsPreAuditSendV2PreAuditMaterialsType   ToolsPreAuditSendV2PreAuditMaterialsType = "IMG"
	SITE_ToolsPreAuditSendV2PreAuditMaterialsType  ToolsPreAuditSendV2PreAuditMaterialsType = "SITE"
	TITLE_ToolsPreAuditSendV2PreAuditMaterialsType ToolsPreAuditSendV2PreAuditMaterialsType = "TITLE"
	VIDEO_ToolsPreAuditSendV2PreAuditMaterialsType ToolsPreAuditSendV2PreAuditMaterialsType = "VIDEO"
)

// All allowed values of ToolsPreAuditSendV2PreAuditMaterialsType enum
var AllowedToolsPreAuditSendV2PreAuditMaterialsTypeEnumValues = []ToolsPreAuditSendV2PreAuditMaterialsType{
	"IMG",
	"SITE",
	"TITLE",
	"VIDEO",
}

// NewToolsPreAuditSendV2PreAuditMaterialsTypeFromValue returns a pointer to a valid ToolsPreAuditSendV2PreAuditMaterialsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPreAuditSendV2PreAuditMaterialsTypeFromValue(v string) (*ToolsPreAuditSendV2PreAuditMaterialsType, error) {
	ev := ToolsPreAuditSendV2PreAuditMaterialsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPreAuditSendV2PreAuditMaterialsType: valid values are %v", v, AllowedToolsPreAuditSendV2PreAuditMaterialsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPreAuditSendV2PreAuditMaterialsType) IsValid() bool {
	for _, existing := range AllowedToolsPreAuditSendV2PreAuditMaterialsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_pre_audit_send_v2_pre_audit_materials_type value
func (v ToolsPreAuditSendV2PreAuditMaterialsType) Ptr() *ToolsPreAuditSendV2PreAuditMaterialsType {
	return &v
}
