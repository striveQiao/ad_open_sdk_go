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

// ToolQuickAppManagementQuickAppGetV2ApiService ToolQuickAppManagementQuickAppGetV2Api service
type ToolQuickAppManagementQuickAppGetV2ApiService service

type ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolQuickAppManagementQuickAppGetV2ApiService
	advertiserId *int64
	status       *[]*ToolQuickAppManagementQuickAppGetV2Status
	page         *int32
	pageSize     *int32
	updateTime   *ToolQuickAppManagementQuickAppGetV2UpdateTime
	searchKey    *string
	quickAppIds  *[]int64
}

// 广告主ID
func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 应用状态:  AUDIT_DOING:审核中  AUDIT_SEND_REJECTED：送审失败 AUDIT_REJECTED:审核失败  AUDIT_ACCEPTED:审核成功 REMOVED：已下架
func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) Status(status []*ToolQuickAppManagementQuickAppGetV2Status) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	r.status = &status
	return r
}

// 页码，默认值为1
func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) Page(page int32) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认值为10，最大不超过100
func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 修改时间
func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) UpdateTime(updateTime ToolQuickAppManagementQuickAppGetV2UpdateTime) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	r.updateTime = &updateTime
	return r
}

// 搜索关键字 快应用名称或者包名的模糊搜索，可以为空，可以传中文
func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) SearchKey(searchKey string) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	r.searchKey = &searchKey
	return r
}

// 快应用ids
func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) QuickAppIds(quickAppIds []int64) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	r.quickAppIds = &quickAppIds
	return r
}

func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) Execute() (*ToolQuickAppManagementQuickAppGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolQuickAppManagementQuickAppGetGet Method for OpenApi2ToolQuickAppManagementQuickAppGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest
*/
func (a *ToolQuickAppManagementQuickAppGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest {
	return &ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolQuickAppManagementQuickAppGetV2Response
func (a *ToolQuickAppManagementQuickAppGetV2ApiService) getExecute(r *ApiOpenApi2ToolQuickAppManagementQuickAppGetGetRequest) (*ToolQuickAppManagementQuickAppGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolQuickAppManagementQuickAppGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tool/quick_app_management/quick_app/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.updateTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "update_time", r.updateTime)
	}
	if r.searchKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_key", r.searchKey)
	}
	if r.quickAppIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quick_app_ids", r.quickAppIds)
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
