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

// ToolsBpAssetManagementShareGetV30ApiService ToolsBpAssetManagementShareGetV30Api service
type ToolsBpAssetManagementShareGetV30ApiService service

type ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest struct {
	ctx            context.Context
	ApiService     *ToolsBpAssetManagementShareGetV30ApiService
	assetType      *ToolsBpAssetManagementShareGetV30AssetType
	instanceId     *int64
	organizationId *int64
	page           *int32
	pageSize       *int32
	shareType      *ToolsBpAssetManagementShareGetV30ShareType
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) AssetType(assetType ToolsBpAssetManagementShareGetV30AssetType) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	r.assetType = &assetType
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) InstanceId(instanceId int64) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	r.instanceId = &instanceId
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) OrganizationId(organizationId int64) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	r.organizationId = &organizationId
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) Page(page int32) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) ShareType(shareType ToolsBpAssetManagementShareGetV30ShareType) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	r.shareType = &shareType
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) Execute() (*ToolsBpAssetManagementShareGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsBpAssetManagementShareGetGet Method for OpenApiV30ToolsBpAssetManagementShareGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest
*/
func (a *ToolsBpAssetManagementShareGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest {
	return &ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsBpAssetManagementShareGetV30Response
func (a *ToolsBpAssetManagementShareGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsBpAssetManagementShareGetGetRequest) (*ToolsBpAssetManagementShareGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsBpAssetManagementShareGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/bp_asset_management/share/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetType == nil {
		return localVarReturnValue, nil, ReportError("assetType is required and must be specified")
	}
	if r.instanceId == nil {
		return localVarReturnValue, nil, ReportError("instanceId is required and must be specified")
	}

	if r.organizationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "organization_id", r.organizationId)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_type", r.assetType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "instance_id", r.instanceId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.shareType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "share_type", r.shareType)
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
