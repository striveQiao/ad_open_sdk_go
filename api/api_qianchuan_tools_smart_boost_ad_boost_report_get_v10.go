/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// QianchuanToolsSmartBoostAdBoostReportGetV10ApiService QianchuanToolsSmartBoostAdBoostReportGetV10Api service
type QianchuanToolsSmartBoostAdBoostReportGetV10ApiService service

type ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest struct {
	ctx             context.Context
	ApiService      *QianchuanToolsSmartBoostAdBoostReportGetV10ApiService
	advertiserId    *int64
	adId            *int64
	adRaiseVersion  *int64
	startTime       *string
	endTime         *string
	page            *int32
	pageSize        *int32
	timeGranularity *string
	filed           *[]string
	orderField      *string
	orderType       *QianchuanToolsSmartBoostAdBoostReportGetV10OrderType
}

// 千川广告主账户ID
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告计划ID
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) AdId(adId int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.adId = &adId
	return r
}

// 起量版本号，通过【获取起量版本信息】接口获取
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) AdRaiseVersion(adRaiseVersion int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.adRaiseVersion = &adRaiseVersion
	return r
}

// 当前起量版本起量开始时间，格式：2021-03-31 16:00:00，通过【获取起量版本信息】接口获取
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) StartTime(startTime string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.startTime = &startTime
	return r
}

// 当前起量版本起量结束时间，格式：2021-03-31 17:00:00，结束时间必须大于开始时间，通过【获取起量版本信息】接口获取
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) EndTime(endTime string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.endTime = &endTime
	return r
}

// 页码，默认值：1
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，允许值：10, 20, 50, 100，默认值：10
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 时间粒度，如果不传，返回查询日期内的聚合数据 允许值: 1、STAT_TIME_GRANULARITY_HOURLY 获取分小时数据
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) TimeGranularity(timeGranularity string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.timeGranularity = &timeGranularity
	return r
}

// 指定需要的指标名称，可参考应答参数返回的消耗指标字段，默认stat_cost，允许值：stat_cost,ctr,convert_cnt,convert_rate,prepay_and_pay_order_roi,show_cnt,click_cnt,pay_order_amount_gmv
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) Filed(filed []string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.filed = &filed
	return r
}

// 排序字段，默认不传为stat_time_hour，允许值：stat_cost,ctr,convert_cnt,convert_rate,prepay_and_pay_order_roi,show_cnt,click_cnt,pay_order_amount_gmv,stat_time_hour
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.orderField = &orderField
	return r
}

// 排序方式
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) OrderType(orderType QianchuanToolsSmartBoostAdBoostReportGetV10OrderType) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) Execute() (*QianchuanToolsSmartBoostAdBoostReportGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGet Method for OpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest
*/
func (a *QianchuanToolsSmartBoostAdBoostReportGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest {
	return &ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsSmartBoostAdBoostReportGetV10Response
func (a *QianchuanToolsSmartBoostAdBoostReportGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostReportGetGetRequest) (*QianchuanToolsSmartBoostAdBoostReportGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanToolsSmartBoostAdBoostReportGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/report/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.adId == nil {
		return localVarReturnValue, nil, ReportError("adId is required and must be specified")
	}
	if r.adRaiseVersion == nil {
		return localVarReturnValue, nil, ReportError("adRaiseVersion is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_raise_version", r.adRaiseVersion)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	if r.timeGranularity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_granularity", r.timeGranularity)
	}
	if r.filed != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filed", r.filed)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
