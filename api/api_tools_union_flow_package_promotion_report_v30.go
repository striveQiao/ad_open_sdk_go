/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
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

// ToolsUnionFlowPackagePromotionReportV30ApiService ToolsUnionFlowPackagePromotionReportV30Api service
type ToolsUnionFlowPackagePromotionReportV30ApiService service

type ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsUnionFlowPackagePromotionReportV30ApiService
	advertiserId *int64
	filter       *ToolsUnionFlowPackagePromotionReportV30Filter
	orderField   *string
	orderType    *ToolsUnionFlowPackagePromotionReportV30OrderType
	page         *int64
	pageSize     *int64
}

// 广告主ID
func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 筛选条件
func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) Filter(filter ToolsUnionFlowPackagePromotionReportV30Filter) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	r.filter = &filter
	return r
}

// 排序字段，所有的统计指标均可参与排序 默认按rit排序
func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) OrderField(orderField string) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	r.orderField = &orderField
	return r
}

// 排序方式 枚举值：ASC（升序），DESC（降序） 默认：DESC
func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) OrderType(orderType ToolsUnionFlowPackagePromotionReportV30OrderType) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	r.orderType = &orderType
	return r
}

// 页数 默认值: 1
func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) Page(page int64) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	r.page = &page
	return r
}

// 页面大小,限制范围[1, 100] 默认值: 10
func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) PageSize(pageSize int64) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) Execute() (*ToolsUnionFlowPackagePromotionReportV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsUnionFlowPackagePromotionReportGet Method for OpenApiV30ToolsUnionFlowPackagePromotionReportGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest
*/
func (a *ToolsUnionFlowPackagePromotionReportV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest {
	return &ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsUnionFlowPackagePromotionReportV30Response
func (a *ToolsUnionFlowPackagePromotionReportV30ApiService) getExecute(r *ApiOpenApiV30ToolsUnionFlowPackagePromotionReportGetRequest) (*ToolsUnionFlowPackagePromotionReportV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsUnionFlowPackagePromotionReportV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/union/flow_package/promotion/report/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.filter == nil {
		return localVarReturnValue, nil, ReportError("filter is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
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
