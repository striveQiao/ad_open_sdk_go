/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAndroidBasicPackageUpdateV2RequestFilesInner struct for ToolsAppManagementAndroidBasicPackageUpdateV2RequestFilesInner
type ToolsAppManagementAndroidBasicPackageUpdateV2RequestFilesInner struct {
	FileTag  *ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTag `json:"file_tag,omitempty"`
	FileType ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType `json:"file_type"`
	// 文件上传id
	UploadId int64 `json:"upload_id"`
}
