/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorUpdateV30RequestAnchorInfoGameAnchor 游戏锚点，当需要更新的锚点类型为anchor_type=APP_GAME: 应用下载-游戏锚点时必填
type NativeAnchorUpdateV30RequestAnchorInfoGameAnchor struct {
	// 游戏图片比例枚举，100：尺寸为2：1的横图，200：尺寸为3：5的竖图
	AnchorImageMode *int64 `json:"anchor_image_mode,omitempty"`
	// 安卓锚点入口标题字段，长度1～12
	AndroidAnchorTitle *string `json:"android_anchor_title,omitempty"`
	// 安卓下载链接，game_type为微信小游戏场景下不用传入
	AndroidDownloadUrl *string `json:"android_download_url,omitempty"`
	// 游戏图片，图片个数 3～8
	AppImages []*NativeAnchorUpdateV30RequestAnchorInfoGameAnchorAppImagesInner `json:"app_images,omitempty"`
	//
	AppOpenUrl *string `json:"app_open_url,omitempty"`
	// 游戏标签列表，每个标签长度1～6，标签个数1～2
	AppTags []string `json:"app_tags,omitempty"`
	// 游戏福利，可选择YES启用、NO不启用
	GameBonus *bool `json:"game_bonus,omitempty"`
	// 游戏特色，长度 1～45
	GameCharatoristic *string `json:"game_charatoristic,omitempty"`
	// 游戏简介，长度 1～45
	GameDescription *string `json:"game_description,omitempty"`
	// 游戏礼包列表，当game_bonus为YES时必填，数量限制0-3
	GamePackageList []*NativeAnchorUpdateV30RequestAnchorInfoGameAnchorGamePackageListInner `json:"game_package_list,omitempty"`
	GameType        *NativeAnchorUpdateV30AnchorInfoGameAnchorGameType                      `json:"game_type,omitempty"`
	// 引导文案，长度1～15
	GuideText *string `json:"guide_text,omitempty"`
	// 锚点头部图片list，推荐尺寸为2：1的横图
	HeadImageList []*NativeAnchorUpdateV30RequestAnchorInfoGameAnchorHeadImageListInner `json:"head_image_list,omitempty"`
	// 小游戏图片list，当前锚点类型且game_type为MICRO_GAME必填，比例为1：1，最多一个
	IconImages []*NativeAnchorUpdateV30RequestAnchorInfoGameAnchorIconImagesInner `json:"icon_images,omitempty"`
	// 微信小游戏ID，当前锚点类型且game_type为MICRO_GAME必填
	InstanceId *int64 `json:"instance_id,omitempty"`
	// iOS锚点入口标题字段，长度1～12
	IosAnchorTitle *string `json:"ios_anchor_title,omitempty"`
	// iOS下载链接，game_type为微信小游戏场景下不用传入
	IosDownloadUrl *string `json:"ios_download_url,omitempty"`
	// 其他说明，长度 1～200
	OtherDescription *string `json:"other_description,omitempty"`
	// 微信小游戏/小程序路径参数
	PathParam *string `json:"path_param,omitempty"`
	// 配置平台，game_type为微信小游戏场景下不用传入
	PlatformType *int64 `json:"platform_type,omitempty"`
}
