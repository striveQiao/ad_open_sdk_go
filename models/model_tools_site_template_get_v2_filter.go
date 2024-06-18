/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateGetV2Filter
type ToolsSiteTemplateGetV2Filter struct {
	// 模板名称关键词
	KeywordOfName *string `json:"keyword_of_name,omitempty"`
	// 站点ID，可通过[【橙子建站】（https://www.chengzijianzhan.com/）]平台或[【获取橙子建站站点列表】](https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710620579852#item-link-%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80)获取
	SiteIds []int64 `json:"site_ids,omitempty"`
	// 模板ID数组，【基于站点创建模板】接口创建的模板ID
	TemplateIds []int64 `json:"template_ids,omitempty"`
}
