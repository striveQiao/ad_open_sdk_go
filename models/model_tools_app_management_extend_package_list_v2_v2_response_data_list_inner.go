/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageListV2V2ResponseDataListInner struct for ToolsAppManagementExtendPackageListV2V2ResponseDataListInner
type ToolsAppManagementExtendPackageListV2V2ResponseDataListInner struct {
	//
	ChannelId *string `json:"channel_id,omitempty"`
	//
	DownloadUrl *string `json:"download_url,omitempty"`
	//
	Md5 *string `json:"md5,omitempty"`
	//
	PackageId *string `json:"package_id,omitempty"`
	//
	Reason *string `json:"reason,omitempty"`
	//
	Remark *string                                                `json:"remark,omitempty"`
	Status *ToolsAppManagementExtendPackageListV2V2DataListStatus `json:"status,omitempty"`
	//
	UpdateTime *string `json:"update_time,omitempty"`
	//
	VersionName *string `json:"version_name,omitempty"`
}
