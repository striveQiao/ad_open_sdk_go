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

// QianchuanReportLongTransferOrderGetV10ApiService QianchuanReportLongTransferOrderGetV10Api service
type QianchuanReportLongTransferOrderGetV10ApiService service

type ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest struct {
	ctx                   context.Context
	ApiService            *QianchuanReportLongTransferOrderGetV10ApiService
	advertiserId          *int64
	startDate             *string
	endDate               *string
	optimizationTimeScope *QianchuanReportLongTransferOrderGetV10OptimizationTimeScope
	filtering             *QianchuanReportLongTransferOrderGetV10Filtering
	marketingGoal         *QianchuanReportLongTransferOrderGetV10MarketingGoal
	orderType             *QianchuanReportLongTransferOrderGetV10OrderType
	page                  *int32
	pageSize              *int32
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 开始时间，格式 2021-04-05
func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) StartDate(startDate string) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间，格式 2021-04-05，时间跨度不能超过90天
func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) EndDate(endDate string) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) OptimizationTimeScope(optimizationTimeScope QianchuanReportLongTransferOrderGetV10OptimizationTimeScope) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.optimizationTimeScope = &optimizationTimeScope
	return r
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) Filtering(filtering QianchuanReportLongTransferOrderGetV10Filtering) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) MarketingGoal(marketingGoal QianchuanReportLongTransferOrderGetV10MarketingGoal) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) OrderType(orderType QianchuanReportLongTransferOrderGetV10OrderType) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) Execute() (*QianchuanReportLongTransferOrderGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanReportLongTransferOrderGetGet Method for OpenApiV10QianchuanReportLongTransferOrderGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest
*/
func (a *QianchuanReportLongTransferOrderGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest {
	return &ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanReportLongTransferOrderGetV10Response
func (a *QianchuanReportLongTransferOrderGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanReportLongTransferOrderGetGetRequest) (*QianchuanReportLongTransferOrderGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanReportLongTransferOrderGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/report/long_transfer/order/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}
	if r.optimizationTimeScope == nil {
		return localVarReturnValue, nil, ReportError("optimizationTimeScope is required and must be specified")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "optimization_time_scope", r.optimizationTimeScope)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
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
