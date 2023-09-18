/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
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

// BusinessPlatformCompanyAccountGetV30ApiService BusinessPlatformCompanyAccountGetV30Api service
type BusinessPlatformCompanyAccountGetV30ApiService service

type ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest struct {
	ctx            context.Context
	ApiService     *BusinessPlatformCompanyAccountGetV30ApiService
	organizationId *int64
	companyId      *int64
	accountType    *[]*BusinessPlatformCompanyAccountGetV30AccountType
	page           *int32
	pageSize       *int32
}

func (r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) OrganizationId(organizationId int64) *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest {
	r.organizationId = &organizationId
	return r
}

func (r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) CompanyId(companyId int64) *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest {
	r.companyId = &companyId
	return r
}

func (r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) AccountType(accountType []*BusinessPlatformCompanyAccountGetV30AccountType) *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest {
	r.accountType = &accountType
	return r
}

func (r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) Page(page int32) *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) PageSize(pageSize int32) *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) Execute() (*BusinessPlatformCompanyAccountGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) WithLog(enable bool) *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BusinessPlatformCompanyAccountGetGet Method for OpenApiV30BusinessPlatformCompanyAccountGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest
*/
func (a *BusinessPlatformCompanyAccountGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest {
	return &ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BusinessPlatformCompanyAccountGetV30Response
func (a *BusinessPlatformCompanyAccountGetV30ApiService) getExecute(r *ApiOpenApiV30BusinessPlatformCompanyAccountGetGetRequest) (*BusinessPlatformCompanyAccountGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *BusinessPlatformCompanyAccountGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/business_platform/company_account/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.organizationId == nil {
		return localVarReturnValue, nil, ReportError("organizationId is required and must be specified")
	}
	if r.companyId == nil {
		return localVarReturnValue, nil, ReportError("companyId is required and must be specified")
	}
	if r.accountType == nil {
		return localVarReturnValue, nil, ReportError("accountType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "organization_id", r.organizationId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "company_id", r.companyId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
