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

// ToolsAipThirdSiteGetV2ApiService ToolsAipThirdSiteGetV2Api service
type ToolsAipThirdSiteGetV2ApiService service

type ApiOpenApi2ToolsAipThirdSiteGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAipThirdSiteGetV2ApiService
	advertiserId *int64
	siteId       *int64
}

// 广告主id
func (r *ApiOpenApi2ToolsAipThirdSiteGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAipThirdSiteGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 站点id，暂不支持批量查询
func (r *ApiOpenApi2ToolsAipThirdSiteGetGetRequest) SiteId(siteId int64) *ApiOpenApi2ToolsAipThirdSiteGetGetRequest {
	r.siteId = &siteId
	return r
}

func (r *ApiOpenApi2ToolsAipThirdSiteGetGetRequest) Execute() (*ToolsAipThirdSiteGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAipThirdSiteGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAipThirdSiteGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAipThirdSiteGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAipThirdSiteGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAipThirdSiteGetGet Method for OpenApi2ToolsAipThirdSiteGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAipThirdSiteGetGetRequest
*/
func (a *ToolsAipThirdSiteGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAipThirdSiteGetGetRequest {
	return &ApiOpenApi2ToolsAipThirdSiteGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAipThirdSiteGetV2Response
func (a *ToolsAipThirdSiteGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAipThirdSiteGetGetRequest) (*ToolsAipThirdSiteGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAipThirdSiteGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/aip_third_site/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.siteId == nil {
		return localVarReturnValue, nil, ReportError("siteId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "site_id", r.siteId)
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
