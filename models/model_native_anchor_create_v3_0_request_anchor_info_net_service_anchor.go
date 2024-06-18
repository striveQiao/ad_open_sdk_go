/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30RequestAnchorInfoNetServiceAnchor 网服下载锚点
type NativeAnchorCreateV30RequestAnchorInfoNetServiceAnchor struct {
	// APP图片比例，100：尺寸为2：1的横图，200：尺寸为3：5的竖图
	AnchorImageMode *int64 `json:"anchor_image_mode,omitempty"`
	// 安卓锚点入口标题字段
	AndroidAnchorTitle *string `json:"android_anchor_title,omitempty"`
	// 安卓下载链接
	AndroidDownloadUrl *string `json:"android_download_url,omitempty"`
	// APP详情，1～100
	AppDescription *string `json:"app_description,omitempty"`
	// APP图片，图片个数 3～8
	AppImages []*NativeAnchorCreateV30RequestAnchorInfoNetServiceAnchorAppImagesInner `json:"app_images,omitempty"`
	// app调起链接
	AppOpenUrl *string `json:"app_open_url,omitempty"`
	// APP标签列表，每个标签长度：1～4；标签个数：1～3
	AppTags []string `json:"app_tags,omitempty"`
	// 引导文案，长度：1～15
	GuideText *string `json:"guide_text,omitempty"`
	// 锚点头部图片list，推荐尺寸为2：1的横图
	HeadImageList []*NativeAnchorCreateV30RequestAnchorInfoNetServiceAnchorHeadImageListInner `json:"head_image_list,omitempty"`
	// 应用icon图片，比例1:1，大小不超过1M，图片个数为1。
	IconImages []*NativeAnchorCreateV30RequestAnchorInfoNetServiceAnchorIconImagesInner `json:"icon_images,omitempty"`
	// 微信小程序资产id
	InstanceId *int64 `json:"instance_id,omitempty"`
	// iOS 锚点入口标题字段
	IosAnchorTitle *string `json:"ios_anchor_title,omitempty"`
	// iOS下载链接
	IosDownloadUrl *string                                                        `json:"ios_download_url,omitempty"`
	NetServiceType *NativeAnchorCreateV30AnchorInfoNetServiceAnchorNetServiceType `json:"net_service_type,omitempty"`
	// 小说章节
	NovelChapter *string `json:"novel_chapter,omitempty"`
	// 路径参数，可选填。如要跳转指定页面，需填写
	PathParam *string `json:"path_param,omitempty"`
	// 配置平台（1:不限,2:安卓,3:iOS）不限：安卓下载链接和iOS下载链接必填；安卓：安卓下载链接必填，iOS下载链接不填写；iOS：iOS下载链接必填
	PlatformType *int64 `json:"platform_type,omitempty"`
	// 快应用资产id
	QuickAppId *int64 `json:"quick_app_id,omitempty"`
	// 微信外跳链接
	WechatExternalUrl *string `json:"wechat_external_url,omitempty"`
	// 微信号码包id或者企业微信号码包id
	WechatPackageId *int64 `json:"wechat_package_id,omitempty"`
}
