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

// ToolsAppManagementExtendPackageListV2V2ApiService ToolsAppManagementExtendPackageListV2V2Api service
type ToolsAppManagementExtendPackageListV2V2ApiService service

type ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest struct {
	ctx         context.Context
	ApiService  *ToolsAppManagementExtendPackageListV2V2ApiService
	accountId   *int64
	accountType *ToolsAppManagementExtendPackageListV2V2AccountType
	packageId   *string
	page        *int32
	pageSize    *int32
	filtering   *ToolsAppManagementExtendPackageListV2V2Filtering
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) AccountId(accountId int64) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	r.accountId = &accountId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) AccountType(accountType ToolsAppManagementExtendPackageListV2V2AccountType) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	r.accountType = &accountType
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) PackageId(packageId string) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	r.packageId = &packageId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) Page(page int32) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) Filtering(filtering ToolsAppManagementExtendPackageListV2V2Filtering) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) Execute() (*ToolsAppManagementExtendPackageListV2V2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementExtendPackageListV2Get Method for OpenApi2ToolsAppManagementExtendPackageListV2Get

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest
*/
func (a *ToolsAppManagementExtendPackageListV2V2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest {
	return &ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementExtendPackageListV2V2Response
func (a *ToolsAppManagementExtendPackageListV2V2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementExtendPackageListV2GetRequest) (*ToolsAppManagementExtendPackageListV2V2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAppManagementExtendPackageListV2V2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/extend_package/list_v2/"

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
	if r.packageId == nil {
		return localVarReturnValue, nil, ReportError("packageId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "package_id", r.packageId)
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
