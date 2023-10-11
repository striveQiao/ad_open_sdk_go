/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.10
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

// QianchuanFinanceDetailGetV10ApiService QianchuanFinanceDetailGetV10Api service
type QianchuanFinanceDetailGetV10ApiService service

type ApiOpenApiV10QianchuanFinanceDetailGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanFinanceDetailGetV10ApiService
	advertiserId *int64
	startDate    *string
	endDate      *string
	type_        *QianchuanFinanceDetailGetV10Type
	page         *int32
	pageSize     *int32
}

// 广告主ID
func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 开始时间，格式 2021-04-05
func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) StartDate(startDate string) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	r.startDate = &startDate
	return r
}

// 结束时间，格式 2021-04-05，时间跨度不能超过1年
func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) EndDate(endDate string) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	r.endDate = &endDate
	return r
}

// 资金池类型
func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) Type_(type_ QianchuanFinanceDetailGetV10Type) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	r.type_ = &type_
	return r
}

// 页码. 默认值: 1
func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) Page(page int32) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	r.page = &page
	return r
}

// 页面数据量. 默认值: 10 上限：200
func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) Execute() (*QianchuanFinanceDetailGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanFinanceDetailGetGet Method for OpenApiV10QianchuanFinanceDetailGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanFinanceDetailGetGetRequest
*/
func (a *QianchuanFinanceDetailGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest {
	return &ApiOpenApiV10QianchuanFinanceDetailGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanFinanceDetailGetV10Response
func (a *QianchuanFinanceDetailGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanFinanceDetailGetGetRequest) (*QianchuanFinanceDetailGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanFinanceDetailGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/finance/detail/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_)
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
