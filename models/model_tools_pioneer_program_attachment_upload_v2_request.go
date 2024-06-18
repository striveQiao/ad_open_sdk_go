/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPioneerProgramAttachmentUploadV2Request struct for ToolsPioneerProgramAttachmentUploadV2Request
type ToolsPioneerProgramAttachmentUploadV2Request struct {
	// 即应用ID，开发者创建应用的唯一标识，可通过应用管理查看
	AppId        int64                                             `json:"app_id"`
	DataFileType ToolsPioneerProgramAttachmentUploadV2DataFileType `json:"data_file_type"`
	// 是否为调试模式，调试模式下数据不会被最终记录，默认为False
	DebugMode *bool `json:"debug_mode,omitempty"`
	// 文件数据，binary格式
	FileData *FormFileInfo `json:"file_data"`
	// 当前文件序号，从1开始，最大值为file_total_num
	FileIndex int64 `json:"file_index"`
	// p_date - platform - data_file_type维度下总文件数，如“2022-06-12”日“巨量引擎”平台的“后端投放数据”共计3份，则file_total_num为3
	FileTotalNum int64 `json:"file_total_num"`
	// 数据日期，格式为yyyy-MM-dd
	PDate    string                                        `json:"p_date"`
	Platform ToolsPioneerProgramAttachmentUploadV2Platform `json:"platform"`
}
