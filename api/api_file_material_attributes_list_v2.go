/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
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

// FileMaterialAttributesListV2ApiService FileMaterialAttributesListV2Api service
type FileMaterialAttributesListV2ApiService service

type ApiOpenApi2FileMaterialAttributesListGetRequest struct {
	ctx                         context.Context
	ApiService                  *FileMaterialAttributesListV2ApiService
	accountId                   *int64
	accountType                 *FileMaterialAttributesListV2AccountType
	pageSize                    *int32
	page                        *int32
	filtering                   *FileMaterialAttributesListV2Filtering
	returnLowqualitySuggestions *bool
}

// 各平台账户id，需同时选择account_type类型 需传入「获取已授权账户」接口返回的advertiser_id
func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) AccountId(accountId int64) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	r.accountId = &accountId
	return r
}

// 账户类型，允许值：
func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) AccountType(accountType FileMaterialAttributesListV2AccountType) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	r.accountType = &accountType
	return r
}

// 页面数据量，默认值: 100，page*page_size 最大1000
func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) PageSize(pageSize int32) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	r.pageSize = &pageSize
	return r
}

// 页码，默认值: 1，page*page_size 最大1000
func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) Page(page int32) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	r.page = &page
	return r
}

// 过滤条件
func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) Filtering(filtering FileMaterialAttributesListV2Filtering) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) ReturnLowqualitySuggestions(returnLowqualitySuggestions bool) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	r.returnLowqualitySuggestions = &returnLowqualitySuggestions
	return r
}

func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) Execute() (*FileMaterialAttributesListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileMaterialAttributesListGetRequest) WithLog(enable bool) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileMaterialAttributesListGet Method for OpenApi2FileMaterialAttributesListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileMaterialAttributesListGetRequest
*/
func (a *FileMaterialAttributesListV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileMaterialAttributesListGetRequest {
	return &ApiOpenApi2FileMaterialAttributesListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileMaterialAttributesListV2Response
func (a *FileMaterialAttributesListV2ApiService) getExecute(r *ApiOpenApi2FileMaterialAttributesListGetRequest) (*FileMaterialAttributesListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileMaterialAttributesListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/material_attributes/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, ReportError("accountId is required and must be specified")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 1 {
		return localVarReturnValue, nil, ReportError("pageSize must be greater than 1")
	}
	if *r.pageSize > 1000 {
		return localVarReturnValue, nil, ReportError("pageSize must be less than 1000")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if *r.page > 1000 {
		return localVarReturnValue, nil, ReportError("page must be less than 1000")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.returnLowqualitySuggestions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "return_lowquality_suggestions", r.returnLowqualitySuggestions)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
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
