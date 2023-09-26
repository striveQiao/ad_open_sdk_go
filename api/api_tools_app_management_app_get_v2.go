/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
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

// ToolsAppManagementAppGetV2ApiService ToolsAppManagementAppGetV2Api service
type ToolsAppManagementAppGetV2ApiService service

type ApiOpenApi2ToolsAppManagementAppGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAppManagementAppGetV2ApiService
	advertiserId *int64
	page         *int32
	pageSize     *int32
	searchKey    *string
	searchType   *ToolsAppManagementAppGetV2SearchType
	status       *ToolsAppManagementAppGetV2Status
	publishTime  *ToolsAppManagementAppGetV2PublishTime
	createTime   *ToolsAppManagementAppGetV2CreateTime
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) Page(page int32) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) SearchKey(searchKey string) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.searchKey = &searchKey
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) SearchType(searchType ToolsAppManagementAppGetV2SearchType) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.searchType = &searchType
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) Status(status ToolsAppManagementAppGetV2Status) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.status = &status
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) PublishTime(publishTime ToolsAppManagementAppGetV2PublishTime) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.publishTime = &publishTime
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) CreateTime(createTime ToolsAppManagementAppGetV2CreateTime) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.createTime = &createTime
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) Execute() (*ToolsAppManagementAppGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementAppGetGet Method for OpenApi2ToolsAppManagementAppGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementAppGetGetRequest
*/
func (a *ToolsAppManagementAppGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppManagementAppGetGetRequest {
	return &ApiOpenApi2ToolsAppManagementAppGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementAppGetV2Response
func (a *ToolsAppManagementAppGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementAppGetGetRequest) (*ToolsAppManagementAppGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAppManagementAppGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/app/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.searchKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_key", r.searchKey)
	}
	if r.searchType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_type", r.searchType)
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
	}
	if r.publishTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "publish_time", r.publishTime)
	}
	if r.createTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_time", r.createTime)
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
