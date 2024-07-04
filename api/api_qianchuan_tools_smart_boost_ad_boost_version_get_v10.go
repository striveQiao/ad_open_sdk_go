/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.11
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

// QianchuanToolsSmartBoostAdBoostVersionGetV10ApiService QianchuanToolsSmartBoostAdBoostVersionGetV10Api service
type QianchuanToolsSmartBoostAdBoostVersionGetV10ApiService service

type ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanToolsSmartBoostAdBoostVersionGetV10ApiService
	advertiserId *int64
	adId         *int64
	cursor       *int64
	count        *int64
}

// 千川广告主账户ID
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告计划ID
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest) AdId(adId int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest {
	r.adId = &adId
	return r
}

// 页码游标值
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest) Cursor(cursor int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest {
	r.cursor = &cursor
	return r
}

// 页面大小
func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest) Count(count int64) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest) Execute() (*QianchuanToolsSmartBoostAdBoostVersionGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGet Method for OpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest
*/
func (a *QianchuanToolsSmartBoostAdBoostVersionGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest {
	return &ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsSmartBoostAdBoostVersionGetV10Response
func (a *QianchuanToolsSmartBoostAdBoostVersionGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostVersionGetGetRequest) (*QianchuanToolsSmartBoostAdBoostVersionGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanToolsSmartBoostAdBoostVersionGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/version/get/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
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
