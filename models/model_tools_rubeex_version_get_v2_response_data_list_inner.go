/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRubeexVersionGetV2ResponseDataListInner struct for ToolsRubeexVersionGetV2ResponseDataListInner
type ToolsRubeexVersionGetV2ResponseDataListInner struct {
	// data_md5指rubeex工具产出的试玩素材作品版本，一个rubeex试玩素材可能有多个作品版本
	DataMd5 *string `json:"data_md5,omitempty"`
	// 首次创建时间
	FirstLaunchDate *string `json:"first_launch_date,omitempty"`
	// 版本对应作品名称
	Name *string `json:"name,omitempty"`
	// 素材url
	PlayableUrl *string `json:"playable_url,omitempty"`
}
