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

// BusinessPlatformPartnerOrganizationListV2ApiService BusinessPlatformPartnerOrganizationListV2Api service
type BusinessPlatformPartnerOrganizationListV2ApiService service

type ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest struct {
	ctx            context.Context
	ApiService     *BusinessPlatformPartnerOrganizationListV2ApiService
	organizationId *int64
	page           *int32
	pageSize       *int32
	filtering      *BusinessPlatformPartnerOrganizationListV2Filtering
}

// 巨量纵横组织id
func (r *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest) OrganizationId(organizationId int64) *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest {
	r.organizationId = &organizationId
	return r
}

// 页码，默认值为1
func (r *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest) Page(page int32) *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认值为10，最大不超过200
func (r *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest) PageSize(pageSize int32) *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest {
	r.pageSize = &pageSize
	return r
}

// 过滤条件
func (r *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest) Filtering(filtering BusinessPlatformPartnerOrganizationListV2Filtering) *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest) Execute() (*BusinessPlatformPartnerOrganizationListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest) AccessToken(accessToken string) *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest) WithLog(enable bool) *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2BusinessPlatformPartnerOrganizationListGet Method for OpenApi2BusinessPlatformPartnerOrganizationListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest
*/
func (a *BusinessPlatformPartnerOrganizationListV2ApiService) Get(ctx context.Context) *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest {
	return &ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BusinessPlatformPartnerOrganizationListV2Response
func (a *BusinessPlatformPartnerOrganizationListV2ApiService) getExecute(r *ApiOpenApi2BusinessPlatformPartnerOrganizationListGetRequest) (*BusinessPlatformPartnerOrganizationListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BusinessPlatformPartnerOrganizationListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/business_platform/partner_organization/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationId == nil {
		return localVarReturnValue, nil, ReportError("organizationId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "organization_id", r.organizationId)
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
