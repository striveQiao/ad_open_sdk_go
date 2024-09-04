/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskCreateV30Request struct for StardeliveryTaskCreateV30Request
type StardeliveryTaskCreateV30Request struct {
	// 广告主账户ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AuthorSubmitFrequency *int32 `json:"author_submit_frequency,omitempty"`
	// 达人侧任务名称，1-50个字符，达人浏览任务时会看到此名称,建议言简意赅、清晰易懂
	AuthorTaskName string `json:"author_task_name"`
	// 请传入一级行业分类ID，请先调用「获取行业列表」接口获取行业ID 当前支持： - 游戏（一级行业id = 1913） - 零售（一级行业id = 1914） - 文体娱乐（一级行业id = 1921） - 传媒资讯（一级行业id = 1922） - 工具类软件（一级行业id = 1928） - 社交通讯（一级行业id= 1935）
	FirstIndustryId int32 `json:"first_industry_id"`
	// 字节小程序链接，当锚点类型为字节小程序时传入有效 - 传参要求：   - 链接中包含的小程序appid必须与锚点内的字节小程序appid一致，否则会报错   - 建议通过「获取字节小程序/小游戏详情内容」接口查询锚点所选字节小程序下可用的链接（对应接口应答参数中的link字段）
	MicroAppLink *string `json:"micro_app_link,omitempty"`
	// 请传入二级行业分类ID，请先调用「获取行业列表」接口获取行业ID 必须传入以下一级行业类目中的二级行业类目ID - 游戏（一级行业id = 1913） - 零售（一级行业id = 1914） - 文体娱乐（一级行业id = 1921） - 传媒资讯（一级行业id = 1922） - 工具类软件（一级行业id = 1928） - 社交通讯（一级行业id= 1935）
	SecondIndustryId int32 `json:"second_industry_id"`
	// 素材出价，只允许整数，每条视频最低需要付给达人的底价价格，该出价影响达人接单的积极性 不同素材类型组合下的出价要求不同，最小可出价金额请调用「获取星广任务金额配置条件」接口查询，最高10000元
	StarMaterialBid float64 `json:"star_material_bid"`
	// 素材一级类目，字段类型是number【特别注意】 任务创建后不允许修改，请前往「获取星广任务可选素材类型」接口获取star_material_first_type（本质是一个id，而不是name）
	StarMaterialFirstType int32 `json:"star_material_first_type"`
	// 素材二级类目，任务创建后不允许修改，需调用「获取星广任务可选素材类型」接口查询一级分类下可选的二级分类名称
	StarMaterialSecondType int32 `json:"star_material_second_type"`
	// 原生锚点id，选择锚点后用于生成组件，并用于达人流量分发，锚点的后续更改不影响组件设置。 - 请先调用「获取原生锚点列表」接口查询审核已通过的锚点，仅支持游戏、网服、电商3类锚点 - 不符合要求接口会报错，传入要求：   1. 仅支持游戏、电商、网服3类锚点   2. 仅支持传入通过审核的锚点素材（status = AUDIT_SUCCESS审核通过）
	StarTaskAnchorId   int64                                             `json:"star_task_anchor_id"`
	StarTaskAnchorType StardeliveryTaskCreateV30StarTaskAnchorType       `json:"star_task_anchor_type"`
	StarTaskAssetInfo  StardeliveryTaskCreateV30RequestStarTaskAssetInfo `json:"star_task_asset_info"`
	// 任务出价（元），最多允许2位小数，用于CPA达人流量计费，设置后不支持修改
	StarTaskBid float64 `json:"star_task_bid"`
	// 任务预算（元），必须是整数。您只能使用广告账户内非赠款下单。最低现金10000，最高1000000元。任务提交后预算可增加、不可减少，任务预算金额影响达人派单总量。
	StarTaskBudget                float64                                                       `json:"star_task_budget"`
	StarTaskExternalAction        StardeliveryTaskCreateV30StarTaskExternalAction               `json:"star_task_external_action"`
	StarTaskMaterialsRequirements StardeliveryTaskCreateV30RequestStarTaskMaterialsRequirements `json:"star_task_materials_requirements"`
	// 任务名称，1-50个字符，本名称仅用于方便自主快速查找和定位任务
	StarTaskName string `json:"star_task_name"`
	// 任务结束时间为投稿截止时间。  - 格式：2024-01-01 精确到天，任务结束时间可调用更新任务接口延长 - 传参要求：   - 任务开始日期与任务结束日期间隔时长≥7天且≤365天，否则报错   - 若任务投稿截止时间+14天不在客户合同有效期内，则不可提交任务
	StarTaskPostEndTime string `json:"star_task_post_end_time"`
	// 任务开始时间，必须填写今天及以后的日期，格式：2024-01-01 精确到天，任务起始时间不允许修改
	StarTaskStartTime string `json:"star_task_start_time"`
	// 任务图标uri，将作为任务头像,展示在达人任务卡片上 上传要求：长宽比为1:1、格式为JPG或JPEG且文件大小小于1M的图片,要求像素为108*108 可先通过「上传广告图片」接口获取图片的uri（即该接口中的id）
	TaskAvatarId    string                                          `json:"task_avatar_id"`
	TaskContactInfo StardeliveryTaskCreateV30RequestTaskContactInfo `json:"task_contact_info"`
	// 任务头图uri，将作为任务头像,展示在达人任务卡片上 长宽比为2:1、格式为JPG或JPEG且文件大小小于5M的图片,建议像素为1125*563 可先通过「上传广告图片」接口获取图片的uri（即该接口中的id）
	TaskHeadImageId *string `json:"task_head_image_id,omitempty"`
	// 产品介绍，1-1000个字符，推广对象的介绍,比如游戏行业,可以简述游戏亮点,游戏属于什么类型,玩法、场景、角色介绍
	TaskProductIntroduction string `json:"task_product_introduction"`
	// 产品名称，1-40个字符，任务推广的产品名称,将展示在达人任务卡片上
	TaskProductName string `json:"task_product_name"`
}
