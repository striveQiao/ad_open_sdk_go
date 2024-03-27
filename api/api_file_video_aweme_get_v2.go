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

// FileVideoAwemeGetV2ApiService FileVideoAwemeGetV2Api service
type FileVideoAwemeGetV2ApiService service

type ApiOpenApi2FileVideoAwemeGetGetRequest struct {
	ctx          context.Context
	ApiService   *FileVideoAwemeGetV2ApiService
	advertiserId *int64
	awemeId      *string
	filtering    *FileVideoAwemeGetV2Filtering
	page         *int64
	pageSize     *int64
	cursor       *string
	count        *int64
}

// 广告主ID
func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 抖音号
func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) AwemeId(awemeId string) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) Filtering(filtering FileVideoAwemeGetV2Filtering) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	r.filtering = &filtering
	return r
}

// 深分页页码，从1开始
func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) Page(page int64) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	r.page = &page
	return r
}

// 深分页每页查询大小
func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 游标参数，初始值为0，当该字段不为空时，将以游标查询为准
func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) Cursor(cursor string) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	r.cursor = &cursor
	return r
}

// 游标查询每页查询大小
func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) Count(count int64) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) Execute() (*FileVideoAwemeGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileVideoAwemeGetGetRequest) WithLog(enable bool) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileVideoAwemeGetGet Method for OpenApi2FileVideoAwemeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileVideoAwemeGetGetRequest
*/
func (a *FileVideoAwemeGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileVideoAwemeGetGetRequest {
	return &ApiOpenApi2FileVideoAwemeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileVideoAwemeGetV2Response
func (a *FileVideoAwemeGetV2ApiService) getExecute(r *ApiOpenApi2FileVideoAwemeGetGetRequest) (*FileVideoAwemeGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileVideoAwemeGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/video/aweme/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
