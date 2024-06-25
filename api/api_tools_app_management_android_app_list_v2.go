/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.9
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

// ToolsAppManagementAndroidAppListV2ApiService ToolsAppManagementAndroidAppListV2Api service
type ToolsAppManagementAndroidAppListV2ApiService service

type ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest struct {
	ctx         context.Context
	ApiService  *ToolsAppManagementAndroidAppListV2ApiService
	accountId   *int64
	accountType *ToolsAppManagementAndroidAppListV2AccountType
	page        *int32
	pageSize    *int32
	filtering   *ToolsAppManagementAndroidAppListV2Filtering
}

func (r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) AccountId(accountId int64) *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest {
	r.accountId = &accountId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) AccountType(accountType ToolsAppManagementAndroidAppListV2AccountType) *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest {
	r.accountType = &accountType
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) Page(page int32) *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) Filtering(filtering ToolsAppManagementAndroidAppListV2Filtering) *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) Execute() (*ToolsAppManagementAndroidAppListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementAndroidAppListGet Method for OpenApi2ToolsAppManagementAndroidAppListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest
*/
func (a *ToolsAppManagementAndroidAppListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest {
	return &ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementAndroidAppListV2Response
func (a *ToolsAppManagementAndroidAppListV2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementAndroidAppListGetRequest) (*ToolsAppManagementAndroidAppListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAppManagementAndroidAppListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/android_app/list/"

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

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
