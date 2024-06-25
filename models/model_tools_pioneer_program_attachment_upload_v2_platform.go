/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPioneerProgramAttachmentUploadV2Platform
type ToolsPioneerProgramAttachmentUploadV2Platform string

// List of tools_pioneer_program_attachment_upload_v2_platform
const (
	KUAI_SHOU_ToolsPioneerProgramAttachmentUploadV2Platform    ToolsPioneerProgramAttachmentUploadV2Platform = "KUAI_SHOU"
	OCEAN_ENGINE_ToolsPioneerProgramAttachmentUploadV2Platform ToolsPioneerProgramAttachmentUploadV2Platform = "OCEAN_ENGINE"
	TENCENT_ToolsPioneerProgramAttachmentUploadV2Platform      ToolsPioneerProgramAttachmentUploadV2Platform = "TENCENT"
)

// All allowed values of ToolsPioneerProgramAttachmentUploadV2Platform enum
var AllowedToolsPioneerProgramAttachmentUploadV2PlatformEnumValues = []ToolsPioneerProgramAttachmentUploadV2Platform{
	"KUAI_SHOU",
	"OCEAN_ENGINE",
	"TENCENT",
}

// NewToolsPioneerProgramAttachmentUploadV2PlatformFromValue returns a pointer to a valid ToolsPioneerProgramAttachmentUploadV2Platform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPioneerProgramAttachmentUploadV2PlatformFromValue(v string) (*ToolsPioneerProgramAttachmentUploadV2Platform, error) {
	ev := ToolsPioneerProgramAttachmentUploadV2Platform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPioneerProgramAttachmentUploadV2Platform: valid values are %v", v, AllowedToolsPioneerProgramAttachmentUploadV2PlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPioneerProgramAttachmentUploadV2Platform) IsValid() bool {
	for _, existing := range AllowedToolsPioneerProgramAttachmentUploadV2PlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_pioneer_program_attachment_upload_v2_platform value
func (v ToolsPioneerProgramAttachmentUploadV2Platform) Ptr() *ToolsPioneerProgramAttachmentUploadV2Platform {
	return &v
}
