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

// CreativeGetV2ApiService CreativeGetV2Api service
type CreativeGetV2ApiService service

type ApiOpenApi2CreativeGetGetRequest struct {
	ctx          context.Context
	ApiService   *CreativeGetV2ApiService
	advertiserId *int64
	filtering    *CreativeGetV2Filtering
	fields       *[]string
	page         *int64
	pageSize     *int64
	cursor       *int64
	count        *int64
}

func (r *ApiOpenApi2CreativeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CreativeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤条件，若此字段不传，或传空则视为无限制条件
func (r *ApiOpenApi2CreativeGetGetRequest) Filtering(filtering CreativeGetV2Filtering) *ApiOpenApi2CreativeGetGetRequest {
	r.filtering = &filtering
	return r
}

// 查询字段集合, 如果指定, 则返回结果数组中, 每个元素是包含所查询字段的字典，默认全部指定 允许值: \&quot;creative_id\&quot;, \&quot;ad_id\&quot;, \&quot;advertiser_id\&quot;, \&quot;status\&quot;,\&quot;opt_status\&quot;, \&quot;image_mode\&quot;, \&quot;title\&quot;, \&quot;creative_word_ids\&quot;,\&quot;third_party_id\&quot;, \&quot;image_ids\&quot;, \&quot;image_id\&quot;, \&quot;video_id\&quot;,\&quot;materials\&quot;
func (r *ApiOpenApi2CreativeGetGetRequest) Fields(fields []string) *ApiOpenApi2CreativeGetGetRequest {
	r.fields = &fields
	return r
}

// 页数默认值: 1，page范围为[1,99999]，2022年1月6号生效。 默认走page翻页
func (r *ApiOpenApi2CreativeGetGetRequest) Page(page int64) *ApiOpenApi2CreativeGetGetRequest {
	r.page = &page
	return r
}

// 页面大小默认值:10，page_size范围为[1,100]，2022年1月6号生效。默认走page翻页
func (r *ApiOpenApi2CreativeGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2CreativeGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 页码游标值，第一次拉取，无需入参 同时传入时，cursor优先级大于page 注：page+page_size与cursor+count为两种分页方式 cursor+count适用于获取数据记录数≥10000的场景
func (r *ApiOpenApi2CreativeGetGetRequest) Cursor(cursor int64) *ApiOpenApi2CreativeGetGetRequest {
	r.cursor = &cursor
	return r
}

// 页面数据量 注：page+page_size与cursor+count为两种分页方式 cursor+count适用于获取数据记录数≥10000的场景
func (r *ApiOpenApi2CreativeGetGetRequest) Count(count int64) *ApiOpenApi2CreativeGetGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApi2CreativeGetGetRequest) Execute() (*CreativeGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CreativeGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeGetGetRequest) WithLog(enable bool) *ApiOpenApi2CreativeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeGetGet Method for OpenApi2CreativeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeGetGetRequest
*/
func (a *CreativeGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2CreativeGetGetRequest {
	return &ApiOpenApi2CreativeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeGetV2Response
func (a *CreativeGetV2ApiService) getExecute(r *ApiOpenApi2CreativeGetGetRequest) (*CreativeGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CreativeGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
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
